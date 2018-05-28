# CRD Validation

This tool is to inline all references for Kubernetes API specifications, which is a workaround in order to avoid [kubernetes/kubernetes#54579][].

## How to use

This tool is built for [kubeflow][] community while it also works for other CRDs by design.

### Clone

```bash
git clone https://github.com/gaocegege/crd-validation
mv ./crd-validation $GOPATH/src/github.com/kubeflow
cd $GOPATH/src/github.com/kubeflow/crd-validation
```

### Build

```bash
go build .
```

### Run

```bash
./crd-validation -c ./crd-validation.yaml tfjob
```

## Acknowledgements

The work is inspired by [ant31/crd-validation][].

[ant31/crd-validation]: https://github.com/ant31/crd-validation
[kubernetes/kubernetes#54579]: https://github.com/kubernetes/kubernetes/issues/54579#issuecomment-370372942
[kubeflow]: https://github.com/kubeflow
