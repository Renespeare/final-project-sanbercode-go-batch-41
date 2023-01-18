package controllers

import (
	"net/http"
	"final-project/database"
	"final-project/repositories"
	categoryModel "final-project/models/category"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	err, categories := repositories.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"data": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCategoryDetail(c *gin.Context) {
	var (
		result gin.H
	)

	var category categoryModel.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)

	err, categories := repositories.GetCategoryDetail(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"data": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context)  {
	var category categoryModel.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "body not found",
		})
		return
	}

	err = repositories.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Insert Category",
	})
}

func UpdateCategory(c *gin.Context)  {
	var category categoryModel.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "body not found",
		})
		return
	}

	category.ID = int64(id)

	err = repositories.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update Category",
	})
}

func DeleteCategory(c *gin.Context)  {
	var category categoryModel.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = repositories.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Category",
	})
}