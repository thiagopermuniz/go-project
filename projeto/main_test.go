package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	r.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	return r
}

func TestMain_Post(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/data", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"post\"", w.Body.String())
}

func TestMain_Get(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/data", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"get\"", w.Body.String())
}
