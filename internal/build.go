package internal

import (
	"strconv"
	"strings"

	"github.com/RedL0tus/pkl-reader-kustomize/internal/msg"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func parameterFromBool(b *bool) *string {
	if b == nil {
		return nil
	}

	val := strconv.FormatBool(*b)
	return &val
}

func formatParameters(p map[string]*string) string {
	params := make([]string, 0, len(p))
	for k, v := range p {
		if v == nil {
			continue
		}
		params = append(params, k+"="+*v)
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + strings.Join(params, "&")
}

func (r kustomizeReader) build(req msg.Build) ([]byte, error) {
	options := krusty.MakeDefaultOptions()
	k := krusty.MakeKustomizer(options)

	target := req.GetPath()

	// Prepend repository if set
	if req.GetRepository() != nil {
		target = *req.GetRepository() + "//" + target
	}

	params := map[string]*string{
		"submodules": parameterFromBool(req.GetSubmodules()),
		"ref":        req.GetRef(),
	}

	target += formatParameters(params)

	resMap, err := k.Run(filesys.MakeFsOnDisk(), target)

	if err != nil {
		return nil, err
	}

	return resMap.AsYaml()
}
