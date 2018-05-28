package crd

import (
	"github.com/kubeflow/crd-validation/pkg/utils"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

type Generator interface {
	Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition
}

type Exporter struct {
	outputDir string
}

func (e Exporter) Export(final *apiextensions.CustomResourceDefinition) string {
	utils.MarshallCrd(final, "yaml")
	return ""
}
