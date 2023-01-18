package controllers

import (
	"final-project/database"
	articleModel "final-project/models/article"
	"final-project/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArticle(c *gin.Context)  {

	var (
		result gin.H
	)

	err, articles := repositories.GetAllArticle(database.DbConnection)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"data": articles,
		}
	}
	
	c.JSON(http.StatusOK, result)
}

func GetArticleDetail(c *gin.Context) {
	var (
		result gin.H
	)

	var article articleModel.Article
	id, _ := strconv.Atoi(c.Param("id"))

	article.ID = int64(id)

	err, articles := repositories.GetArticleDetail(database.DbConnection, article)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"data": articles,
		}
	}

	c.JSON(http.StatusOK, result)
}


func InsertArticle(c *gin.Context)  {

	var (
		result gin.H
	)

	var article articleModel.Article

	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "body not found",
		})
		return
	}
	userId := c.MustGet("id")

	err = repositories.InsertArticle(database.DbConnection, article, userId.(float64))
	if err != nil {
		result = gin.H{
			"error": "data category not found",
		}
	} else {
		result = gin.H{
			"message": "Success Insert Article",
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateArticle(c *gin.Context)  {

	var (
		result gin.H
	)

	var article articleModel.Article
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "body not found",
		})
		return
	}

	article.ID = int64(id)

	userId := c.MustGet("id")
	err = repositories.UpdateArticle(database.DbConnection, article, userId.(float64))
	if err != nil {
		result = gin.H{
			"error": "data category not found",
		}
	} else {
		result = gin.H{
			"message": "Success Update Article",
		}
	}

	c.JSON(http.StatusOK, result)	
}

func DeleteArticle(c *gin.Context)  {

	var (
		result gin.H
	)

	var article articleModel.Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	article.ID = int64(id)
	
	userId := c.MustGet("id")
	err = repositories.DeleteArticle(database.DbConnection, article, userId.(float64))
	if err != nil {
		result = gin.H{
			"error": err.Error(),
		} 
	} else {
		result = gin.H{
			"message": "Success Delete Article",
		}
	}

	c.JSON(http.StatusOK, result)
}