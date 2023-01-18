package controllers

import (
	"final-project/database"
	commentModel "final-project/models/comment"
	"final-project/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllComment(c *gin.Context)  {

	var (
		result gin.H
	)

	var comment commentModel.Comment
	article_id, _ := strconv.Atoi(c.Query("article_id")) 

	comment.Article_id = int(article_id)

	err, comments := repositories.GetAllComment(database.DbConnection, comment)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"data": comments,
		}
	}

	c.JSON(http.StatusOK, result)
	
}

func GetCommentDetail(c *gin.Context) {
	var (
		result gin.H
	)

	var comment commentModel.Comment
	id, _ := strconv.Atoi(c.Param("id"))

	comment.ID = int64(id)

	err, articles := repositories.GetCommentDetail(database.DbConnection, comment)

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

func InsertComment(c *gin.Context)  {

	var (
		result gin.H
	)

	var comment commentModel.Comment
	article_id, _ := strconv.Atoi(c.Query("article_id")) 

	comment.Article_id = int(article_id)

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}

	userId := c.MustGet("id")
	err = repositories.InsertComment(database.DbConnection, comment, userId.(float64))
	if err != nil {
		result = gin.H{
			"error": "data article not found",
		}
	} else {
		result = gin.H{
			"message": "Success Insert Comment",
		}
	}

	c.JSON(http.StatusOK, result)	
}

func DeleteComment(c *gin.Context)  {

	var (
		result gin.H
	)

	var comment commentModel.Comment
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	comment.ID = int64(id)
	
	userId := c.MustGet("id")
	err = repositories.DeleteComment(database.DbConnection, comment, userId.(float64))
	if err != nil {
		result = gin.H{
			"error": err.Error(),
		} 
	} else {
		result = gin.H{
			"message": "Success Delete Comment",
		}
	}

	c.JSON(http.StatusOK, result)
}
