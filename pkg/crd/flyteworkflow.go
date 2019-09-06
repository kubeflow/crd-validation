package crd

import (
	"fmt"
	"github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/lyft/flytepropeller/pkg/crd-validation/pkg/crd/exporter"
	"github.com/lyft/flytepropeller/pkg/crd-validation/pkg/utils"
)

const (
	// CRDName is the name for FlyteWorkflow.
	CRDNameFlyteWorkflow = "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1.FlyteWorkflow"

	generatedFileFlyteWorkflow = "flyteworkflow-crd-v1alpha1.yaml"
)

// TFJobGenerator is the type for TFJob CRD generator.
type FlyteWorkflowGenerator struct {
	*exporter.Exporter
}

// NewTFJobGenerator creates a new TFJob CRD generator.
func NewFlyteWorkflowGenerator(outputDir string) *FlyteWorkflowGenerator {
	return &FlyteWorkflowGenerator{
		Exporter: exporter.New(outputDir, generatedFileFlyteWorkflow),
	}
}

// Generate generates the crd.
func (t FlyteWorkflowGenerator) Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition {
	fmt.Println("Generating validation")
	original.Spec.Validation = utils.GetCustomResourceValidation(CRDNameFlyteWorkflow, v1alpha1.GetOpenAPIDefinitions)
	return original
}