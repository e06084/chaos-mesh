// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +chaos-mesh:base

// CustomizeChaos is the Schema for the timechaos API
type CustomizeChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a time chaos experiment
	Spec CustomizeChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the time chaos experiment
	Status CustomizeChaosStatus `json:"status"`
}

// CustomizeChaosSpec defines the desired state of CustomizeChaos
type CustomizeChaosSpec struct {
	// Mode defines the mode to run chaos action.
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent
	// +kubebuilder:validation:Enum=one;all;fixed;fixed-percent;random-max-percent
	Mode PodMode `json:"mode"`

	// Value is required when the mode is set to `FixedPodMode` / `FixedPercentPodMod` / `RandomMaxPercentPodMod`.
	// If `FixedPodMode`, provide an integer of pods to do chaos action.
	// If `FixedPercentPodMod`, provide a number from 0-100 to specify the percent of pods the server can do chaos action.
	// If `RandomMaxPercentPodMod`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	// +optional
	Value string `json:"value"`

	// Selector is used to select pods that are used to inject chaos action.
	Selector SelectorSpec `json:"selector"`

	// ContainerName indicates the name of affected container.
	// If not set, all containers will be injected
	// +optional
	ContainerNames []string `json:"containerNames,omitempty"`

	// Duration represents the duration of the chaos action
	Duration *string `json:"duration,omitempty"`

	// Scheduler defines some schedule rules to control the running time of the chaos experiment about time.
	Scheduler *SchedulerSpec `json:"scheduler,omitempty"`

	// +optional
	ApplyCommand []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`

	// +optional
	ApplyArgs []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`

	// +optional
	RecoverCommand []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`

	// +optional
	RecoverArgs []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
}

// GetSelector is a getter for Selector (for implementing SelectSpec)
func (in *CustomizeChaosSpec) GetSelector() SelectorSpec {
	return in.Selector
}

// GetMode is a getter for Mode (for implementing SelectSpec)
func (in *CustomizeChaosSpec) GetMode() PodMode {
	return in.Mode
}

// GetValue is a getter for Value (for implementing SelectSpec)
func (in *CustomizeChaosSpec) GetValue() string {
	return in.Value
}

// CustomizeChaosStatus defines the observed state of CustomizeChaos
type CustomizeChaosStatus struct {
	ChaosStatus `json:",inline"`
}
