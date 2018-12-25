package main

import (
	"errors"
	"flag"
	"github.com/krmahadevan/di/dig"
	"github.com/krmahadevan/di/manual"
	"github.com/krmahadevan/di/wire"
)

func main() {
	impl := flag.String("type", "wire",
		"Which type of Dependency Injection to use (manual/dig/wire")
	switch *impl {
	case "manual":
		manual.Main()
	case "dig":
		dig.Main()
	case "wire":
		wire.Main()
	default:
		panic(errors.New("unknown Dependency Injection model chosen"))

	}
}
