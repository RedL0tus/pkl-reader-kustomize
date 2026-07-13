package main

import (
	_ "github.com/RedL0tus/pkl-reader-kustomize"
	"github.com/RedL0tus/pkl-reader-kustomize/internal"
	shared "github.com/apple/pkl-readers/shared/go"
)

var (
	Version   = "0.1.1"
	_, _, run = shared.New(shared.Spec{
		SchemeSuffix: "kustomize",
		Name:         "pkl-reader-kustomize",
		Short:        "Pkl External Reader for Kustomize manifests",
		Long: `Pkl External Reader for Kustomize manifests.

External Readers: https://pkl-lang.org/main/current/language-reference/index.html#external-readers

CLI configuration:
	--external-resource-reader reader+kustomize=pkl-reader-kustomize

PklProject configuration:
	evaluatorSettings {
		externalResourceReaders {
			["reader+kustomize"] {
				executable = "pkl-reader-kustomize"
			}
		}
	}
`,
		Version: Version,
	}, internal.Run)
)

func main() {
	run()
}
