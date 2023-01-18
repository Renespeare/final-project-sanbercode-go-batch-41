package controllers

import (
	"final-project/database"
	userModel "final-project/models/user"
	"final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var user userModel.Register

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.Register(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Success!"})   
}

func Login(c *gin.Context) {
	
	var user userModel.Login

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	err, token := repositories.Login(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})

}

func Logout(c *gin.Context)  {
	uuid, _ := c.Get("uuid")
	err := repositories.Logout(database.DbConnection, uuid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user has been logout"})
}