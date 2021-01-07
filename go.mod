module github.com/fyuan1316/flagger-operator

go 1.13

replace (
	github.com/fyuan1316/klient v0.0.0-20210106101119-962edf8f6e33 => ./local/klient // indirect
	github.com/fyuan1316/operatorlib v0.0.0-20210106132823-879be4d125b8 => ./local/operatorlib
)

require (
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/coreos/prometheus-operator v0.41.0 // indirect
	github.com/fyuan1316/asm-operator v0.0.0-20201228103053-53e0299f8171
	github.com/fyuan1316/klient v0.0.0-20210106101119-962edf8f6e33 // indirect
	github.com/fyuan1316/operatorlib v0.0.0-20210106132823-879be4d125b8
	github.com/go-logr/logr v0.2.1
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	helm.sh/helm/v3 v3.4.2 // indirect
	k8s.io/api v0.19.4 // indirect
	k8s.io/apiextensions-apiserver v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/yaml v1.2.0
)
