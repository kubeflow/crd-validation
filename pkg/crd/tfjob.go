package crd

import (
	tfv1alpha2 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	log "github.com/sirupsen/logrus"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/kubeflow/crd-validation/pkg/crd/exporter"
	"github.com/kubeflow/crd-validation/pkg/utils"
)

const (
	// CRDName is the name for TFJob.
	CRDName = "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2.TFJob"

	generatedFile = "tfjob-crd-v1alpha2.yaml"
)

// TFJobGenerator is the type for TFJob CRD generator.
type TFJobGenerator struct {
	*exporter.Exporter
}

// NewTFJobGenerator creates a new TFJob CRD generator.
func NewTFJobGenerator(outputDir string) *TFJobGenerator {
	return &TFJobGenerator{
		Exporter: exporter.New(outputDir, generatedFile),
	}
}

// Generate generates the crd.
func (t TFJobGenerator) Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition {
	log.Println("Generating validation for TFJob")
	original.Spec.Validation = utils.GetCustomResourceValidation(CRDName, tfv1alpha2.GetOpenAPIDefinitions)
	return original
}
