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

package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	extensionsobj "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func MarshallCrd(crd *extensionsobj.CustomResourceDefinition, outputFormat string) {
	jsonBytes, err := json.MarshalIndent(crd, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}

	if outputFormat == "json" {
		os.Stdout.Write(jsonBytes)
	} else {
		yamlBytes, err := yaml.JSONToYAML(jsonBytes)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write([]byte("---\n"))
		os.Stdout.Write(yamlBytes)
	}
}
