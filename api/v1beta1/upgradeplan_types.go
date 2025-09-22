/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UpgradePlanConditionType defines UpgradePlan's condition
type UpgradePlanConditionType string

const (
	// UpgradePlanAvailable means the UpgradePlan is validated and ready to be kicked off.
	UpgradePlanAvailable UpgradePlanConditionType = "Available"
	// UpgradePlanProgressing means the cluster is currently applying the UpgradePlan.
	UpgradePlanProgressing UpgradePlanConditionType = "Progressing"
	// UpgradePlanDegraded means the cluster encounters issues executing the UpgradePlan.
	UpgradePlanDegraded UpgradePlanConditionType = "Degraded"
)

// UpgradePlanPhase defines what overall phase UpgradePlan is in
type UpgradePlanPhase string

const (
	UpgradePlanPhaseInit           UpgradePlanPhase = "Init"
	UpgradePlanPhaseImagePreload   UpgradePlanPhase = "ImagePreload"
	UpgradePlanPhaseClusterUpgrade UpgradePlanPhase = "ClusterUpgrade"
	UpgradePlanPhaseNodeUpgrade    UpgradePlanPhase = "NodeUpgrade"
	UpgradePlanPhaseCleanup        UpgradePlanPhase = "Cleanup"
	UpgradePlanPhaseSucceeded      UpgradePlanPhase = "Succeeded"
	UpgradePlanPhaseFailed         UpgradePlanPhase = "Failed"
)

type UpgradePlanPhaseTransitionTimestamp struct {
	Phase                    UpgradePlanPhase `json:"phase,omitempty"`
	PhaseTransitionTimestamp metav1.Time      `json:"phaseTransitionTimestamp,omitempty"`
}

// UpgradePlanSpec defines the desired state of UpgradePlan
type UpgradePlanSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	//  version refers to the corresponding version resource in the same namespace.
	Version string `json:"version"`

	// upgrade can be specified to opt for any other specific upgrade image. If not provided, the version resource name is used.
	// For instance, specifying "dev" for the field can go for the "rancher/harvester-upgrade:dev" image.
	// +optional
	Upgrade *string `json:"upgrade,omitempty"`

	// force indicates the UpgradePlan will be forcibly applied, ignoring any pre-upgrade check failures. Default to "false".
	// +optional
	Force *bool `json:"force,omitempty"`

	// mode represents the manipulative style of the UpgradePlan. Can be either of "automatic" or "interactive". Default to "automatic".
	// +optional
	Mode *string `json:"mode,omitempty"`
}

// UpgradePlanStatus defines the observed state of UpgradePlan.
type UpgradePlanStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the UpgradePlan resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// phase shows what overall phase the UpgradePlan resource is in
	Phase UpgradePlanPhase `json:"phase,omitempty"`

	// phaseTransitionTimestamp is the timestamp of when the last phase change occurred.
	// listType=atomic
	// +optional
	PhaseTransitionTimestamps []UpgradePlanPhaseTransitionTimestamp `json:"phaseTransitionTimestamps,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// UpgradePlan is the Schema for the upgradeplans API
type UpgradePlan struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of UpgradePlan
	// +required
	Spec UpgradePlanSpec `json:"spec"`

	// status defines the observed state of UpgradePlan
	// +optional
	Status UpgradePlanStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// UpgradePlanList contains a list of UpgradePlan
type UpgradePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpgradePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpgradePlan{}, &UpgradePlanList{})
}
