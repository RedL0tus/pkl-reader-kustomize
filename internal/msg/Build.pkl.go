// Code generated from Pkl module `kustomize.kustomize`. DO NOT EDIT.
package msg

type Build interface {
	Request

	GetRepository() *string

	GetPath() string

	GetRef() *string

	GetSubmodules() *bool
}

var _ Build = BuildImpl{}

type BuildImpl struct {
	Kind string `pkl:"kind"`

	Repository *string `pkl:"repository"`

	Path string `pkl:"path"`

	Ref *string `pkl:"ref"`

	Submodules *bool `pkl:"submodules"`
}

func (rcv BuildImpl) GetKind() string {
	return rcv.Kind
}

func (rcv BuildImpl) GetRepository() *string {
	return rcv.Repository
}

func (rcv BuildImpl) GetPath() string {
	return rcv.Path
}

func (rcv BuildImpl) GetRef() *string {
	return rcv.Ref
}

func (rcv BuildImpl) GetSubmodules() *bool {
	return rcv.Submodules
}
