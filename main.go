package main

import (
	"fmt"
	"souschef-api/config" // Importeer de config voor databaseverbinding
	"souschef-api/models" // Importeer de models voor de struct van recepten
	"souschef-api/routes" // Importeer de routes voor de API-endpoints
)

func main() {
	// Verbind de applicatie met de database
	config.ConnectDatabase()

	// Voer database-migratie uit voor de Recipe model
	config.DB.AutoMigrate(&models.Recipe{})

	// Stel de routes in voor de server
	r := routes.SetupRouter()

	// Print een bericht naar de console om aan te geven dat de server draait
	fmt.Println("Server draait op http://localhost:8080")

	// Start de server op poort 8080
	r.Run(":8080")
}
