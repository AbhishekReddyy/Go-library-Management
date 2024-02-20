package routes

import (
	"github.com/AbhishekReddyy/Go-library-Management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/dashboard", controllers.Dashboard())
	incomingroutes.POST("/addbook", controllers.AddBook())
	incomingroutes.POST("/delete", controllers.DeleteBook())
	incomingroutes.POST("/new-user", controllers.NewUser())
	incomingroutes.GET("/search/:id", controllers.SearchUser())

}
