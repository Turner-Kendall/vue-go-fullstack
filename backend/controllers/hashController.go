package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
)

func hashString(i string) string {
	hasher := md5.New()
	hasher.Write([]byte(i))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// Struct to bind JSON input
type HashRequest struct {
	Input string `json:"input"`
}

func GetHash(c *gin.Context) {
	var req HashRequest

	// Try to bind JSON body
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	hash := hashString(req.Input)
	c.JSON(200, gin.H{"hash": hash})
}
