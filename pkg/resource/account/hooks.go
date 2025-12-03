// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package account

import (
	"context"
	"fmt"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/organizations"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/aws-controllers-k8s/organizations-controller/pkg/tags"
)

func (rm *resourceManager) customUpdateAccount(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {

	updated = rm.concreteResource(desired.DeepCopy())

	desiredTags, _ := convertToOrderedACKTags(desired.ko.Spec.Tags)
	latestTags, _ := convertToOrderedACKTags(latest.ko.Spec.Tags)

	if delta.DifferentAt("Spec.Tags") {
		if err := tags.SyncTags(
			ctx, rm.sdkapi, rm.metrics,
			*latest.ko.Status.AccountID,
			desiredTags,
			latestTags,
		); err != nil {
			return nil, err
		}
	}

	newDesired := rm.concreteResource(desired.DeepCopy())
	newDesired.ko.Status = updated.ko.Status
	return newDesired, nil
}

func (rm *resourceManager) fetchCurrentTags(
	ctx context.Context,
	resourceID *string,
) (map[string]string, error) {
	return tags.FetchTags(ctx, rm.sdkapi, rm.metrics, *resourceID)
}

// getAccountID fetches the accountID after Account creation succeeds
// returns requeue error while account creation is in progress
func (rm *resourceManager) getAccountID(ctx context.Context, desired *resource) (accountID *string, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.getAccountID")
	defer func() {
		exit(err)
	}()
	input := &svcsdk.DescribeCreateAccountStatusInput{
		CreateAccountRequestId: desired.ko.Status.CreateAccountRequestID,
	}
	resp, err := rm.sdkapi.DescribeCreateAccountStatus(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCreateAccountStatus", err)
	if err != nil {
		return nil, err
	}

	switch resp.CreateAccountStatus.State {
	case svcsdktypes.CreateAccountStateSucceeded:
		return resp.CreateAccountStatus.AccountId, nil

	case svcsdktypes.CreateAccountStateFailed:
		return nil, errors.NewTerminalError(fmt.Errorf("account creation failed"))

	case svcsdktypes.CreateAccountStateInProgress:
		return nil, requeue.NeededAfter(
			fmt.Errorf("account creation in progress"),
			requeue.DefaultRequeueAfterDuration,
		)

	default:
		return nil, fmt.Errorf("unknown account creation state %s", resp.CreateAccountStatus.State)
	}

}
