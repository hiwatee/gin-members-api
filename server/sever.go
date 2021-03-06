package server

import (
	"members/controllers"
	"members/middleware"
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
	r := Router()
	r.Run()
}

// Router ...
func Router() *gin.Engine {
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
		healthcheck := v1.Group("/healthcheck")
		{
			healthcheck.GET("", controllers.HealthCheckController{}.Index)
		}

		users := v1.Group("/users")
		{
			users.Use(middleware.UserAuth())
			users.GET("", controllers.UserController{}.Index)
			users.GET("/:id", controllers.UserController{}.Show)
			users.PUT("/:id", controllers.UserController{}.Update)
			users.DELETE("/:id", controllers.UserController{}.Delete)
		}

		registration := v1.Group("/registration")
		{
			registration.POST("", controllers.RegistrationController{}.Create)
		}

		login := v1.Group("/login")
		{
			login.POST("", controllers.LoginController{}.Create)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
