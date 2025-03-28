package routes

import (
	"souschef-api/controllers" // Importeer de controllers voor recepten

	"github.com/gin-gonic/gin" // Importeer de Gin framework voor routing
)

// SetupRouter configureert de routes voor de API
func SetupRouter() *gin.Engine {
	r := gin.Default() // Maak een nieuwe Gin router met de standaardinstellingen

	// Stel de verschillende HTTP-methoden in voor de 'recipes' endpoint
	r.GET("/recipes", controllers.GetRecipes)          // Haal alle recepten op
	r.POST("/recipes", controllers.CreateRecipe)       // Voeg een nieuw recept toe
	r.GET("/recipes/:id", controllers.GetRecipeByID)   // Haal een specifiek recept op via ID
	r.DELETE("/recipes/:id", controllers.DeleteRecipe) // Verwijder een recept op basis van ID

	return r // Retourneer de router zodat deze kan worden gebruikt in de main.go
}
