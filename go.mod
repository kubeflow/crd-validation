module github.com/kubeflow/crd-validation

go 1.13

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.19.7
	github.com/kubeflow/common v0.0.0-20190619012831-09e1ac17011c // indirect
	github.com/kubeflow/tf-operator v1.0.0-rc.0.0.20200401015727-95a0f622c13b
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
	k8s.io/apiextensions-apiserver v0.15.10
	k8s.io/apimachinery v0.15.10
	k8s.io/client-go v10.0.0+incompatible // indirect
	k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/kubernetes v1.15.10 // indirect
	k8s.io/sample-controller v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.15.10
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.15.10
	k8s.io/apimachinery => k8s.io/apimachinery v0.15.10
	k8s.io/apiserver => k8s.io/apiserver v0.15.10
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.15.10
	k8s.io/client-go => k8s.io/client-go v0.15.10
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.15.10
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.15.10
	k8s.io/code-generator => k8s.io/code-generator v0.15.10
	k8s.io/component-base => k8s.io/component-base v0.15.10
	k8s.io/cri-api => k8s.io/cri-api v0.15.10
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.15.10
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.15.10
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.15.10
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.15.10
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.15.10
	k8s.io/kubectl => k8s.io/kubectl v0.15.10
	k8s.io/kubelet => k8s.io/kubelet v0.15.10
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.15.10
	k8s.io/metrics => k8s.io/metrics v0.15.10
	k8s.io/node-api => k8s.io/node-api v0.15.10
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.15.10
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.15.10
	k8s.io/sample-controller => k8s.io/sample-controller v0.15.10
)
