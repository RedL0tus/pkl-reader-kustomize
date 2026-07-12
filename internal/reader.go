package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/RedL0tus/pkl-reader-kustomize/internal/msg"
	"github.com/apple/pkl-go/pkl"
	shared "github.com/apple/pkl-readers/shared/go"
)

type Options struct{}

func Run(ctx context.Context, spec shared.Spec, _ *Options) error {
	reader := kustomizeReader{
		Spec: spec,
	}

	return shared.Run(ctx, spec, pkl.WithExternalClientResourceReader(reader))
}

type kustomizeReader struct {
	shared.Spec
}

func (r kustomizeReader) Read(uri url.URL) ([]byte, error) {
	var req msg.Request
	if err := r.DecodeRequest(uri, &req); err != nil {
		return nil, err
	}

	slog.Debug("received request", "kind", req.GetKind())

	switch reqType := req.(type) {
	case msg.Build:
		return r.build(reqType)
	default:
		return nil, fmt.Errorf("unrecognized action '%s'", uri.Host)
	}
}
