package main

import (
	"fmt"
	"go-rest/database"
	"go-rest/routes"
)

func main() {
	database.ConnectToDB()
	fmt.Println("Starting server")
	routes.HandleRequest()
}
