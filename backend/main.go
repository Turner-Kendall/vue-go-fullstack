package main

import (
	"net/http"

	"github.com/Turner-Kendall/gapi/controllers"
	"github.com/Turner-Kendall/gapi/initializers"
	"github.com/Turner-Kendall/gapi/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	err := router.SetTrustedProxies([]string{"192.168.1.25"})
	if err != nil {
		panic(err)
	}

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hey Now!"})
	})
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)
	router.GET("/hash", middleware.CheckAuth, controllers.GetHash)
	router.GET("/users", middleware.CheckAuth, controllers.ListUsers)
	router.DELETE("/user/delete", middleware.CheckAuth, controllers.DeleteUser)
	router.PUT("/user/password", middleware.CheckAuth, controllers.ResetPassword)
	router.Run()
}
