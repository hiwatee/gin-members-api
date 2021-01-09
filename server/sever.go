package server

import (
	"members/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// docs ...
	_ "members/docs"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			ctrl := controllers.UserController{}
			users.GET("", ctrl.Index)
			users.POST("", ctrl.Create)
			users.GET("/:id", ctrl.Show)
			users.PUT("/:id", ctrl.Update)
			users.DELETE("/:id", ctrl.Delete)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
