// migrate/migrate.go

package main

import (
	"github.com/Turner-Kendall/gapi/initializers"
	"github.com/Turner-Kendall/gapi/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
