package wire

import "fmt"

func Main() {
	fmt.Println("Making use of [wire] as Dependency Injection")
	server, err := NewServerInstance()
	if err != nil {
		panic(err)
	}
	server.Run()
}
