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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccountSpec defines the desired state of Account.
//
// Contains information about an Amazon Web Services account that is a member
// of an organization.
type AccountSpec struct {

	// The email address of the owner to assign to the new member account. This
	// email address must not already be associated with another Amazon Web Services
	// account. You must use a valid email address to complete account creation.
	//
	// The rules for a valid email address:
	//
	//   - The address must be a minimum of 6 and a maximum of 64 characters long.
	//
	//   - All characters must be 7-bit ASCII characters.
	//
	//   - There must be one and only one @ symbol, which separates the local name
	//     from the domain name.
	//
	// +kubebuilder:validation:Required
	Email *string `json:"email"`
	// If set to ALLOW, the new account enables IAM users to access account billing
	// information if they have the required permissions. If set to DENY, only the
	// root user of the new account can access account billing information. For
	// more information, see About IAM access to the Billing and Cost Management
	// console (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	//
	// If you don't specify this parameter, the value defaults to ALLOW, and IAM
	// users and roles with the required permissions can access billing information
	// for the new account.
	IAMUserAccessToBilling *string `json:"iamUserAccessToBilling,omitempty"`
	// The friendly name of the member account.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The name of an IAM role that Organizations automatically preconfigures in
	// the new member account. This role trusts the management account, allowing
	// users in the management account to assume the role, as permitted by the management
	// account administrator. The role has administrator permissions in the new
	// member account.
	//
	// If you don't specify this parameter, the role name defaults to OrganizationAccountAccessRole.
	//
	// For more information about how to use this role to access the member account,
	// see the following links:
	//
	//   - Creating the OrganizationAccountAccessRole in an invited member account
	//     (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role)
	//     in the Organizations User Guide
	//
	//   - Steps 2 and 3 in IAM Tutorial: Delegate access across Amazon Web Services
	//     accounts using IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)
	//     in the IAM User Guide
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter. The pattern can include uppercase letters, lowercase letters,
	// digits with no spaces, and any of the following characters: =,.@-
	RoleName *string `json:"roleName,omitempty"`
	// A list of tags that you want to attach to the newly created account. For
	// each tag in the list, you must specify both a tag key and a value. You can
	// set the value to an empty string, but you can't set it to null. For more
	// information about tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)
	// in the Organizations User Guide.
	//
	// If any one of the tags is not valid or if you exceed the maximum allowed
	// number of tags for an account, then the entire request fails and the account
	// is not created.
	Tags []*Tag `json:"tags,omitempty"`
}

// AccountStatus defines the observed state of Account
type AccountStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// If the account was created successfully, the unique identifier (ID) of the
	// new account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an account ID string
	// requires exactly 12 digits.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountID,omitempty"`
	// The date and time that the account was created and the request completed.
	// +kubebuilder:validation:Optional
	CompletedTimestamp *metav1.Time `json:"completedTimestamp,omitempty"`
	// The unique identifier (ID) that references this request. You get this value
	// from the response of the initial CreateAccount request to create the account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a create account
	// request ID string requires "car-" followed by from 8 to 32 lowercase letters
	// or digits.
	// +kubebuilder:validation:Optional
	CreateAccountRequestID *string `json:"createAccountRequestID,omitempty"`
	// If the request failed, a description of the reason for the failure.
	//
	//    * ACCOUNT_LIMIT_EXCEEDED: The account couldn't be created because you
	//    reached the limit on the number of accounts in your organization.
	//
	//    * CONCURRENT_ACCOUNT_MODIFICATION: You already submitted a request with
	//    the same information.
	//
	//    * EMAIL_ALREADY_EXISTS: The account could not be created because another
	//    Amazon Web Services account with that email address already exists.
	//
	//    * FAILED_BUSINESS_VALIDATION: The Amazon Web Services account that owns
	//    your organization failed to receive business license validation.
	//
	//    * GOVCLOUD_ACCOUNT_ALREADY_EXISTS: The account in the Amazon Web Services
	//    GovCloud (US) Region could not be created because this Region already
	//    includes an account with that email address.
	//
	//    * IDENTITY_INVALID_BUSINESS_VALIDATION: The Amazon Web Services account
	//    that owns your organization can't complete business license validation
	//    because it doesn't have valid identity data.
	//
	//    * INVALID_ADDRESS: The account could not be created because the address
	//    you provided is not valid.
	//
	//    * INVALID_EMAIL: The account could not be created because the email address
	//    you provided is not valid.
	//
	//    * INVALID_PAYMENT_INSTRUMENT: The Amazon Web Services account that owns
	//    your organization does not have a supported payment method associated
	//    with the account. Amazon Web Services does not support cards issued by
	//    financial institutions in Russia or Belarus. For more information, see
	//    Managing your Amazon Web Services payments (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-general.html).
	//
	//    * INTERNAL_FAILURE: The account could not be created because of an internal
	//    failure. Try again later. If the problem persists, contact Amazon Web
	//    Services Customer Support.
	//
	//    * MISSING_BUSINESS_VALIDATION: The Amazon Web Services account that owns
	//    your organization has not received Business Validation.
	//
	//    * MISSING_PAYMENT_INSTRUMENT: You must configure the management account
	//    with a valid payment method, such as a credit card.
	//
	//    * PENDING_BUSINESS_VALIDATION: The Amazon Web Services account that owns
	//    your organization is still in the process of completing business license
	//    validation.
	//
	//    * UNKNOWN_BUSINESS_VALIDATION: The Amazon Web Services account that owns
	//    your organization has an unknown issue with business license validation.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// If the account was created successfully, the unique identifier (ID) of the
	// new account in the Amazon Web Services GovCloud (US) Region.
	// +kubebuilder:validation:Optional
	GovCloudAccountID *string `json:"govCloudAccountID,omitempty"`
	// The date and time that the request was made for the account creation.
	// +kubebuilder:validation:Optional
	RequestedTimestamp *metav1.Time `json:"requestedTimestamp,omitempty"`
	// The status of the asynchronous request to create an Amazon Web Services account.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// Account is the Schema for the Accounts API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec,omitempty"`
	Status            AccountStatus `json:"status,omitempty"`
}

// AccountList contains a list of Account
// +kubebuilder:object:root=true
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
