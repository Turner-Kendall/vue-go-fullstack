package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gosimple/slug"

	"github.com/Turner-Kendall/gapi/initializers"
	"github.com/Turner-Kendall/gapi/models"
	"github.com/Turner-Kendall/gapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var regInput models.RegistrationInput

	if err := c.ShouldBindJSON(&regInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if the username is already in use
	var userFound models.User

	initializers.DB.Where("username=?", regInput.Username).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	initializers.DB.Where("email=?", regInput.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	// hash the password + check if the password is valid
	passwordHash, err := utils.ProcessPassword(regInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// first make the username lower case - so our blacklist is more effective.
	username, err := utils.ProcessUsername(regInput.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// then create a new slug
	slug.MaxLength = 50
	slug.CustomSub = utils.ReturnBlackList()
	usernameSlug := slug.Make(username)

	user := models.User{
		Lastname:  regInput.Lastname,
		Firstname: regInput.Firstname,
		Email:     regInput.Email,
		Phone:     regInput.Phone,
		Username:  usernameSlug,
		Password:  string(passwordHash),
		Key:       ksuid.New().String(),
		Provider:  "gÄuth",
		Dob:       regInput.Dob.Time,
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func Login(c *gin.Context) {

	var authInput models.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)
	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// func GetUserProfile(c *gin.Context) {

// 	user, _ := c.Get("currentUser")

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }
