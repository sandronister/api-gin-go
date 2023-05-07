package main

import (
	"api-go-gin/controllers"
	"api-go-gin/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GetTestRoute() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestVerifyStatusCodeGretings(t *testing.T) {
	r := GetTestRoute()
	r.GET("/:name", controllers.Gretting)
	req, _ := http.NewRequest("GET", "/Julio", nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)
	assert.Equal(t, http.StatusOK, result.Code)
	mockResult := `{"message":"Bem vindo Julio tudo bem?"}`
	bodyResult, _ := ioutil.ReadAll(result.Body)
	assert.Equal(t, mockResult, string(bodyResult))
}

func TestStudentList(t *testing.T) {
	database.ConnectDB()
	r := GetTestRoute()
	r.GET("/student", controllers.ListStundent)
	req, _ := http.NewRequest("GET", "/student", nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)
	assert.Equal(t, http.StatusOK, result.Code)
	mockResult := `[{"ID":1,"CreatedAt":"2023-05-06T16:29:30.215971-03:00","UpdatedAt":"2023-05-06T16:29:30.215971-03:00","DeletedAt":null,"name":"Ana","cpf":"3456","rg":"7090"},{"ID":2,"CreatedAt":"2023-05-06T16:36:29.62959-03:00","UpdatedAt":"2023-05-06T16:36:29.62959-03:00","DeletedAt":null,"name":"David","cpf":"12345678903","rg":"123456789"},{"ID":3,"CreatedAt":"2023-05-06T16:36:42.251405-03:00","UpdatedAt":"2023-05-06T16:36:42.251405-03:00","DeletedAt":null,"name":"Jose","cpf":"1234123412341234","rg":"34123412341234"},{"ID":5,"CreatedAt":"2023-05-06T21:48:30.669933-03:00","UpdatedAt":"2023-05-06T21:48:30.669933-03:00","DeletedAt":null,"name":"Jose boni","cpf":"12345678901","rg":"123456789"}]`
	bodyResult, _ := ioutil.ReadAll(result.Body)
	assert.Equal(t, mockResult, string(bodyResult))
}
