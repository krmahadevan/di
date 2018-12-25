package dig

import (
	"fmt"
	"github.com/krmahadevan/di"
	"go.uber.org/dig"
)

func Main() {
	fmt.Println("Making use of [dig] as Dependency Injection")
	container := BuildContainer()
	err := container.Invoke(func(server *di.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(di.NewConfig)
	_ = container.Provide(di.ConnectDatabase)
	_ = container.Provide(di.NewPersonRepository)
	_ = container.Provide(di.NewPersonService)
	_ = container.Provide(di.NewServer)
	return container
}
