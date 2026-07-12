pkl-reader-kustomize
====================

A [Pkl](https://pkl-lang.org) external resource reader for [Kustomize](https://github.com/kubernetes-sigs/kustomize)/[Krusty](https://pkg.go.dev/sigs.k8s.io/kustomize/api/krusty) builds.

Examples
--------

Read manifests: [gateway-api.pkl](examples/gateway-api.pkl).

Generate Pkl modules from CRDs: [gateway-api-crds](examples/gateway-api-crds.pkl)

Install
-------

1. Build and install binary: `go install -v https://github.com/RedL0tus/pkl-reader-kustomize/cmd/pkl-reader-kustomize@v0.1.0`, make sure the binary exists in PATH.
2. Update PklProject:
  - Add `package://pkg.pkl-lang.org/github.com/RedL0tus/pkl-reader-kustomize@v0.1.0` to dependencies.
  - Update external reader settings:
  ```
  evaluatorSettings {
    externalResourceReaders {
      ["reader+kustomize"] {
        executable = "pkl-reader-kustomize"
      }
    }
  }
  ```

Credits
-------

- Official [pkl-readers](https://github.com/apple/pkl-readers), used their shared utilities and boilerplates in this project.
- [k8s.contrib.crd](https://pkl-lang.org/package-docs/pkg.pkl-lang.org/pkl-pantry/k8s.contrib.crd/4.2.2/generate/index.html), for Pkl module generation from CRDs.
