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

// OrganizationalUnitSpec defines the desired state of OrganizationalUnit.
//
// Contains details about an organizational unit (OU). An OU is a container
// of Amazon Web Services accounts within a root of an organization. Policies
// that are attached to an OU apply to all accounts contained in that OU and
// in any child OUs.
type OrganizationalUnitSpec struct {

	// The friendly name to assign to the new OU.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The unique identifier (ID) of the parent root or OU that you want to create
	// the new OU in.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//   - Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//     letters or digits.
	//
	//   - Organizational unit (OU) - A string that begins with "ou-" followed
	//     by from 4 to 32 lowercase letters or digits (the ID of the root that the
	//     OU is in). This string is followed by a second "-" dash and from 8 to
	//     32 additional lowercase letters or digits.
	//
	// +kubebuilder:validation:Required
	ParentID *string `json:"parentID"`
	// A list of tags that you want to attach to the newly created OU. For each
	// tag in the list, you must specify both a tag key and a value. You can set
	// the value to an empty string, but you can't set it to null. For more information
	// about tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)
	// in the Organizations User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed number of
	// tags for an OU, then the entire request fails and the OU is not created.
	Tags []*Tag `json:"tags,omitempty"`
}

// OrganizationalUnitStatus defines the observed state of OrganizationalUnit
type OrganizationalUnitStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The unique identifier (ID) associated with this OU.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational
	// unit ID string requires "ou-" followed by from 4 to 32 lowercase letters
	// or digits (the ID of the root that contains the OU). This string is followed
	// by a second "-" dash and from 8 to 32 additional lowercase letters or digits.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
}

// OrganizationalUnit is the Schema for the OrganizationalUnits API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationalUnitSpec   `json:"spec,omitempty"`
	Status            OrganizationalUnitStatus `json:"status,omitempty"`
}

// OrganizationalUnitList contains a list of OrganizationalUnit
// +kubebuilder:object:root=true
type OrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationalUnit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OrganizationalUnit{}, &OrganizationalUnitList{})
}
