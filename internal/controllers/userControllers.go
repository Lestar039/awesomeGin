package controllers

import (
	"awesomeGin/config"
	"awesomeGin/internal/models"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func encodePassword(s string) string {
	addSalt := s + os.Getenv("AUTH_SALT")
	h := sha1.New()
	h.Write([]byte(addSalt))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func UserCreate(c *gin.Context) {
	var userStruct struct {
		Login    string
		Password string
	}
	c.Bind(&userStruct)

	user := models.User{Login: userStruct.Login, Password: encodePassword(userStruct.Password)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UsersShow(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"user count": len(users),
		"users":      users,
	})
}

func UserShow(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Find(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	var userStruct struct {
		Login    string
		Password string
	}
	c.Bind(&userStruct)

	var user models.User
	id := c.Param("id")
	config.DB.Find(&user, id)

	config.DB.Model(&user).Updates(models.User{
		Login:    userStruct.Login,
		Password: userStruct.Password,
	})

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {

	id := c.Param("id")
	config.DB.Delete(&models.User{}, id)

	c.JSON(http.StatusOK, gin.H{
		"user": "user deleted",
	})
}

func Login(c *gin.Context) {
	var userStruct struct {
		Login    string
		Password string
	}
	c.Bind(&userStruct)
	checkPassword := encodePassword(userStruct.Password)

	var user models.User
	config.DB.Where("login = ? AND password = ?", userStruct.Login, checkPassword).Find(&user)

	if (models.User{}) == user {
		c.Status(401)
		return
	}
	config.DB.Model(&user).Updates(models.User{
		IsActive:  true,
		LastLogin: time.Now(),
	})
	c.JSON(http.StatusOK, gin.H{
		"login": fmt.Sprintf("User %v login", user.Login),
	})
}

func Logout(c *gin.Context) {
	var userStruct struct {
		Login string
	}
	c.Bind(&userStruct)

	var user models.User
	config.DB.Where("login = ?", userStruct.Login).Find(&user)

	if (models.User{}) == user {
		c.Status(401)
		return
	}
	config.DB.Model(&user).Updates(models.User{
		IsActive: false,
	})
	c.JSON(http.StatusOK, gin.H{
		"login": fmt.Sprintf("User %v logout", user.Login),
	})
}
