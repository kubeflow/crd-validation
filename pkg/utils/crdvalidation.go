package utils

import (
	"github.com/go-openapi/spec"
	extensionsobj "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/kube-openapi/pkg/common"
)

// OpenAPIRefCallBack returns a jsonref using the input string without modification
func OpenAPIRefCallBack(name string) spec.Ref {
	return spec.MustCreateRef(name)
}

// GetAPIDefinition is a function returning a map with all Definition
type GetAPIDefinitions func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition

// GetCustomResourceValidation returns the validation definition for a CRD name
func GetCustomResourceValidation(name string, fn func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition) *extensionsobj.CustomResourceValidation {
	openapiSpec := fn(OpenAPIRefCallBack)
	fixKnownTypes(openapiSpec)

	schema := openapiSpec[name].Schema

	ret := &extensionsobj.CustomResourceValidation{
		OpenAPIV3Schema: SchemaPropsToJSONProps(&schema, openapiSpec, true),
	}
	return ret
}

// ref: https://github.com/kubernetes/kubernetes/issues/62329
func fixKnownTypes(openapiSpec map[string]common.OpenAPIDefinition) {
	openapiSpec["k8s.io/apimachinery/pkg/util/intstr.IntOrString"] = common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				AnyOf: []spec.Schema{
					{
						SchemaProps: spec.SchemaProps{
							Type: []string{"string"},
						},
					},
					{
						SchemaProps: spec.SchemaProps{
							Type: []string{"integer"},
						},
					},
				},
			},
		},
	}
}
