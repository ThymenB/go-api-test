package models

import "gorm.io/gorm" // Import de GORM package voor database interactie

// Recipe is de structuur die een recept vertegenwoordigt in de database
type Recipe struct {
	gorm.Model          // Biedt automatisch velden zoals ID, CreatedAt, UpdatedAt en DeletedAt
	Title        string `json:"title"`        // De titel van het recept
	Ingredients  string `json:"ingredients"`  // De ingrediënten van het recept
	Instructions string `json:"instructions"` // De instructies voor het recept
}
