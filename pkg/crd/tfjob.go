package crd

import (
	"fmt"

	"github.com/kubeflow/crd-validation/pkg/crd/exporter"
	"github.com/kubeflow/crd-validation/pkg/utils"
	tfv1alpha2 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	tfv1beta1 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1beta1"
	log "github.com/sirupsen/logrus"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// TFJobGenerator is the type for TFJob CRD generator.
type TFJobGenerator struct {
	*exporter.Exporter
	jobVersion string
}

// NewTFJobGenerator creates a new TFJob CRD generator.
func NewTFJobGenerator(version string) *TFJobGenerator {
	return &TFJobGenerator{
		Exporter:   exporter.New(),
		jobVersion: version,
	}
}

// Generate generates the crd.
func (t *TFJobGenerator) Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition {
	log.Println("Generating validation for TFJob")

	// CRDName is the name for TFJob.
	CRDName := fmt.Sprintf("github.com/kubeflow/tf-operator/pkg/apis/tensorflow/%v.TFJob", t.jobVersion)

	if t.jobVersion == "v1alpha2" {
		original.Spec.Validation = utils.GetCustomResourceValidation(CRDName, tfv1alpha2.GetOpenAPIDefinitions)
	} else if t.jobVersion == "v1beta1" {
		original.Spec.Validation = utils.GetCustomResourceValidation(CRDName, tfv1beta1.GetOpenAPIDefinitions)
	} else {
		log.Errorf("Invalid version of TFJob %s", t.jobVersion)
	}

	return original
}
