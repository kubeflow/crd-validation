// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/spf13/viper"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewCustomResourceDefinition creates a new CRD from the config for the given name.
func NewCustomResourceDefinition(name string) *apiextensions.CustomResourceDefinition {
	crdConfig := viper.Get(name).(map[string]interface{})
	metadata := crdConfig["metadata"].(map[string]interface{})
	spec := crdConfig["spec"].(map[string]interface{})
	names := spec["names"].(map[string]interface{})
	crd := &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: metadata["name"].(string),
			// Labels: metadata["labels"].(map[string]string),
			// Annotations: ,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       crdConfig["kind"].(string),
			APIVersion: crdConfig["apiversion"].(string),
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   spec["group"].(string),
			Version: spec["version"].(string),
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:   names["plural"].(string),
				Singular: names["singular"].(string),
				Kind:     names["kind"].(string),
			},
			Scope: apiextensions.ResourceScope(spec["scope"].(string)),
		},
	}

	return crd
}
