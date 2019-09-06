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
	"github.com/ghodss/yaml"
	"github.com/spf13/viper"

	"io/ioutil"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"log"
	"os"
)

type CrdValidationConfig struct {
	OutputDir string
}

func GetCrdValidationConfig() *CrdValidationConfig {
	var newConfig CrdValidationConfig
	if viper.Get("global") != nil {
		newConfig.OutputDir = viper.Get("global").(map[string]interface{})["output"].(string)
	}
	return &newConfig
}

// NewCustomResourceDefinition creates a new CRD from the config for the given name.
func NewCustomResourceDefinition(baseCrdFile string) *apiextensions.CustomResourceDefinition {
	log.Printf("Reading base CRD from %v\n", baseCrdFile)
	data, err := ioutil.ReadFile(baseCrdFile)
	var baseCrdObj apiextensions.CustomResourceDefinition

	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	err = yaml.Unmarshal([]byte(data), &baseCrdObj)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	return &baseCrdObj
}
