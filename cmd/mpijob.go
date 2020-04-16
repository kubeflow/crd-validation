// Copyright 2020 The Kubeflow Authors
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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kubeflow/crd-validation/pkg/config"
	"github.com/kubeflow/crd-validation/pkg/crd"
)

func init() {
	RootCmd.AddCommand(mpijobCmd)
}

// tfjobCmd represents the tfjob command
var mpijobCmd = &cobra.Command{
	Use:   "mpijob",
	Short: "Generate MPIJob CRD definition",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		generateMPIJob()
	},
}

func generateMPIJob() {
	original := config.NewCustomResourceDefinition("mpijob")
	var outputDir string
	if viper.Get("global") != nil {
		outputDir = viper.Get("global").(map[string]interface{})["output"].(string)
	}
	generator := crd.NewMPIJobGenerator(outputDir)
	final := generator.Generate(original)
	generator.Export(final)
}
