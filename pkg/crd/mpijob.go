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
package crd

import (
	commonv1 "github.com/kubeflow/common/job_controller/api/v1"
	mpiapis "github.com/kubeflow/mpi-operator/pkg/apis/kubeflow/v1alpha2"
	log "github.com/sirupsen/logrus"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/kubeflow/crd-validation/pkg/crd/exporter"
	"github.com/kubeflow/crd-validation/pkg/utils"
)

const (
	// MPICRDName is the name for MPIJob
	MPICRDName = "github.com/kubeflow/mpi-operator/pkg/apis/kubeflow/v1alpha2.MPIJob"

	mpiGeneratedFile = "mpijob-v1alpha2-crd.yaml"
)

// TFJobGenerator is the type for TFJob CRD generator.
type MPIJobGenerator struct {
	*exporter.Exporter
}

// NewTFJobGenerator creates a new TFJob CRD generator.
func NewMPIJobGenerator(outputDir string) *MPIJobGenerator {
	return &MPIJobGenerator{
		Exporter: exporter.New(outputDir, mpiGeneratedFile),
	}
}

// Generate generates the crd.
func (t MPIJobGenerator) Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition {
	log.Println("Generating validation for MPIJob")
	original.Spec.Validation = utils.GetCustomResourceValidation(
		MPICRDName,
		[]utils.GetAPIDefinitions{
			mpiapis.GetOpenAPIDefinitions,
			commonv1.GetOpenAPIDefinitions,
		})
	return original
}
