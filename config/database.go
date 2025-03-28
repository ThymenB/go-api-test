package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Laadt omgevingsvariabelen uit .env bestand
	"gorm.io/driver/postgres"  // PostgreSQL driver voor GORM
	"gorm.io/gorm"             // GORM ORM voor database-interactie
)

var DB *gorm.DB // Global database variable

// ConnectDatabase maakt verbinding met de PostgreSQL database.
func ConnectDatabase() {
	// Laadt het .env bestand om omgevingsvariabelen in te stellen
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // Foutmelding als .env bestand niet gevonden wordt
	}

	// Stel de database-verbindingstring samen met wachtwoord uit de omgevingsvariabele
	dsn := fmt.Sprintf(
		"host=localhost user=postgres password=%s dbname=souschef port=5432 sslmode=disable",
		os.Getenv("DB_PASSWORD"), // Haalt het wachtwoord op uit de omgevingsvariabele
	)

	// Maak verbinding met de database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err) // Foutmelding als de verbinding mislukt
	}

	// Bewaar de databaseverbinding voor verder gebruik in de app
	DB = database
	fmt.Println("Database connected successfully!") // Bevestiging van succesvolle verbinding
}
