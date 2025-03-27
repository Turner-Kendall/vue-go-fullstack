package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func HashString(i string) string {
	hasher := md5.New()             // Create a new MD5 hash generator
	hasher.Write([]byte(i))         // Write the input string as bytes to the hasher
	hash := hasher.Sum(nil)         // Finalize the hash and get the result as a byte slice
	return hex.EncodeToString(hash) // Convert the byte slice to a hexadecimal string
}

func GetHash(c *gin.Context) {
	input := c.Query("input")
	hash := HashString(input)
	c.JSON(200, gin.H{"hash": hash})
}
