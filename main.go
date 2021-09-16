package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/database"
	"github.com/shubhamdixit863/golangApis/routes"
)

func main() {
	app := fiber.New()
	client := database.Connect("Go_db", "mongodb+srv://keshika:Keshika@cluster0.nlwsi.mongodb.net/")
	defer database.Disconnect(client) // Disconnecting once the main finished execution
	routes.RegisterRoutes(app)

	app.Listen(":3001")

}

//calling handles and display function

/* package main

import (
	"log"
	"net/http"

	"github.com/KeshikaGupta20/Goprogm/handlers"
)

func main() {

	http.HandleFunc("/getData", handlers.GetHandler)
	http.HandleFunc("/display", handlers.Display)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error", err)
	}

} */
