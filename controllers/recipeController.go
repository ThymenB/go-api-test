package controllers

import (
	"net/http"
	"souschef-api/config" // Verbindt met de databaseconfiguratie
	"souschef-api/models" // Verbindt met de modelstructuren van de app

	"github.com/gin-gonic/gin" // Gin web framework
)

// GET /recipes - Haal alle recepten op uit de database
func GetRecipes(c *gin.Context) {
	var recipes []models.Recipe
	// Haalt alle recepten op uit de database
	config.DB.Find(&recipes)
	// Stuurt de recepten terug als JSON met een 200 OK status
	c.JSON(http.StatusOK, recipes)
}

// POST /recipes - Voeg een nieuw recept toe aan de database
func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	// Bind de JSON payload naar het recipe model
	if err := c.ShouldBindJSON(&recipe); err != nil {
		// Stuurt foutmelding als de JSON niet correct is
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Slaat het nieuwe recept op in de database
	config.DB.Create(&recipe)
	// Stuurt het nieuwe recept terug met een 201 Created status
	c.JSON(http.StatusCreated, recipe)
}

// GET /recipes/:id - Haal een specifiek recept op via het id
func GetRecipeByID(c *gin.Context) {
	var recipe models.Recipe
	id := c.Param("id") // Haal het id uit de URL-parameter

	// Zoekt het recept op met het opgegeven id
	if err := config.DB.First(&recipe, id).Error; err != nil {
		// Stuurt foutmelding als het recept niet gevonden wordt
		c.JSON(http.StatusNotFound, gin.H{"error": "Recept niet gevonden"})
		return
	}
	// Stuurt het recept terug als JSON met een 200 OK status
	c.JSON(http.StatusOK, recipe)
}

// DELETE /recipes/:id - Verwijder een specifiek recept uit de database
func DeleteRecipe(c *gin.Context) {
	var recipe models.Recipe
	id := c.Param("id") // Haal het id uit de URL-parameter

	// Verwijder het recept met het opgegeven id
	if err := config.DB.Delete(&recipe, id).Error; err != nil {
		// Stuurt foutmelding als het verwijderen mislukt
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Verwijderen mislukt"})
		return
	}
	// Bevestigt dat het recept succesvol verwijderd is
	c.JSON(http.StatusOK, gin.H{"message": "Recept verwijderd"})
}
