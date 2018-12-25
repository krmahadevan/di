package manual

import (
	"fmt"
	"github.com/krmahadevan/di"
)

// The manual way
//
func Main() {
	fmt.Println("Handcrafting Dependency Injection manually")
	config := di.NewConfig()
	fmt.Println("Database to be read from ", config.DatabasePath)

	db, err := di.ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := di.NewPersonRepository(db)
	personService := di.NewPersonService(config, personRepository)
	server := di.NewServer(config, personService)
	server.Run()
}
