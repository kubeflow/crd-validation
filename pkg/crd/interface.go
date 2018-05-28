package crd

import apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

type Generator interface {
	Generate(original *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinition
}
