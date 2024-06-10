/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceDependencySpec defines the desired state of ServiceDependency
type ServiceDependencySpec struct {
	Service      ServiceDetails `json:"service"`
	Dependencies []Dependency   `json:"dependencies"`
}

type ServiceDetails struct {
	DeploymentName string `json:"deploymentName"`
	Namespace      string `json:"namespace"`
}

type Dependency struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// ServiceDependencyStatus defines the observed state of ServiceDependency
type ServiceDependencyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServiceDependency is the Schema for the servicedependencies API
type ServiceDependency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDependencySpec   `json:"spec,omitempty"`
	Status ServiceDependencyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceDependencyList contains a list of ServiceDependency
type ServiceDependencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDependency `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDependency{}, &ServiceDependencyList{})
}
