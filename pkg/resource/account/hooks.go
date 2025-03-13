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

	"github.com/aws-controllers-k8s/organizations-controller/pkg/tags"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
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
