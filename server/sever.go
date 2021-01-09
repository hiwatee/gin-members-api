package server

import (
	"members/controllers"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8000",
			"http://127.0.0.1:8000",
			"http://0.0.0.0:8000",
		},
		AllowMethods: []string{
			"HEAD",
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		MaxAge:           24 * time.Hour,
		AllowCredentials: true,
	}))

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", controllers.UserController{}.Index)
			users.POST("", controllers.UserController{}.Create)
			users.GET("/:id", controllers.UserController{}.Show)
			users.PUT("/:id", controllers.UserController{}.Update)
			users.DELETE("/:id", controllers.UserController{}.Delete)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
