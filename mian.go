package main

import (
	"log"
	"os"

	"github.com/AbhishekReddyy/Go-library-Management/controllers"
	"github.com/AbhishekReddyy/Go-library-Management/database"
	"github.com/AbhishekReddyy/Go-library-Management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.BookData(database.Client, "Books"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	log.Fatal(router.Run(":", port))
}
