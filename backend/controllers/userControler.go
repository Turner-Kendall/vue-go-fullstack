package controllers

import (
	"net/http"

	"github.com/Turner-Kendall/gapi/initializers"
	"github.com/Turner-Kendall/gapi/models"
	"github.com/Turner-Kendall/gapi/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Get the current user from the context
// https://github.com/gin-gonic/gin/blob/master/context.go
func GetUserProfile(c *gin.Context) {

	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	// Cast user to the appropriate type
	currentUser := user.(models.User)

	// Create a UserProfile with the required fields
	userProfile := models.UserProfile{
		ID:       currentUser.ID,
		Username: currentUser.Username,
		Email:    currentUser.Email,
	}

	c.JSON(200, gin.H{
		"user": userProfile,
	})
}

// list all of the users
func ListUsers(c *gin.Context) {
	var users []models.User // Ensure this is a slice of the correct struct type

	// Query the database for all users
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Check if any users were found
	if len(users) == 0 {
		c.JSON(404, gin.H{
			"message": "No users found in the database.",
		})
		return
	}

	// Return the list of users
	c.JSON(200, gin.H{
		"users": users,
	})
}

// delete user by username
func DeleteUser(c *gin.Context) {

	var delInput models.DelInput

	if err := c.ShouldBindJSON(&delInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User

	initializers.DB.Where("username=?", delInput.Username).Find(&userFound)
	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
		return
	}

	initializers.DB.Delete(&userFound)

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// update user password by username
func ResetPassword(c *gin.Context) {
	var resetInput models.ResetInput

	// Parse JSON input
	if err := c.ShouldBindJSON(&resetInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user from the database
	var userFound models.User
	if err := initializers.DB.Where("username = ?", resetInput.Username).First(&userFound).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
		return
	}

	// Debugging logs
	// fmt.Println("Stored hash:", userFound.Password)
	// fmt.Println("Input old password:", resetInput.OldPassword)

	// Compare hashed password with plaintext password
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(resetInput.OldPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// Hash new password
	passwordHash, err := utils.ProcessPassword(resetInput.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update password in the database
	if err := initializers.DB.Model(&userFound).Update("password", passwordHash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}
