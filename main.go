package main

import (
	"fmt"

	"github.com/droveda/api-rest/database"
	"github.com/droveda/api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Starting the server!")
	routes.HandleRequest()
}
