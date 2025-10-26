package routes

import (
	"golang-test/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/addresses", controllers.CreateAddress)
		api.GET("/addresses", controllers.FindAddresses)
		api.GET("/addresses/:id", controllers.FindAddress)
		api.PUT("/addresses/:id", controllers.UpdateAddress)
		api.DELETE("/addresses/:id", controllers.DeleteAddress)
	}
}
