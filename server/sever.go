package server

import (
	"members/controllers"

	"github.com/gin-gonic/gin"

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

	u := r.Group("api/v1/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
