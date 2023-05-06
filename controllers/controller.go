package controllers

import (
	"api-go-gin/database"
	"api-go-gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllStudent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Show",
	})
}

func Gretting(c *gin.Context) {
	var name = c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Bem vindo " + name + " tudo bem?",
	})
}

func CreateStudent(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func ListStundent(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	var id = c.Params.ByName("id")
	var student model.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var id = c.Params.ByName("id")
	var student model.Student
	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, gin.H{})
}
