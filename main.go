package main

import (
	"fmt"
	"go-rest/database"
	"go-rest/routes"
)

func main() {
	// models.Personalities = []models.Personality{
	// 	{Id: 1, Name: "Name 1", History: "History 1"},
	// 	{Id: 2, Name: "Name 2", History: "History 2"},
	// }
	database.ConnectToDB()
	fmt.Println("Starting server")
	routes.HandleRequest()

}
