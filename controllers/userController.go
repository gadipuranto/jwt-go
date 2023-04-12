package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-jwt/configs"
	"go-jwt/domains/models"
	"go-jwt/helpers"
	"net/http"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := configs.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	userData := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&userData)
	} else {
		c.ShouldBind(&userData)
	}

	userData.Password = helpers.HashPass(userData.Password)
	
	err := db.Debug().Create(&userData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        userData.ID,
		"email":     userData.Email,
		"full_name": userData.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := configs.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	fmt.Println(password)

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}