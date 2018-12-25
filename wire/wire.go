//+build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/krmahadevan/di"
)

func NewServerInstance() (*di.Server, error) {
	wire.Build(di.NewConfig,
		di.ConnectDatabase, //Since this can return an error, we need to ensure we return back that same error
		di.NewPersonRepository,
		di.NewPersonService,
		di.NewServer)
	return &di.Server{}, nil
}
