package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"projeto/internal/service"
)

type DataHandler struct {
	service service.DataServiceInterface
}

func NewDataHandler(svc service.DataServiceInterface) *DataHandler {
	return &DataHandler{
		service: svc,
	}
}

func (h *DataHandler) GetHandler(c *gin.Context) {
	key := c.Query("key")
	data, err := h.service.GetServiceData(c.Request.Context(), key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			c.JSON(http.StatusNotFound, gin.H{"message": "key not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	jsonData := map[string]any{}
	_ = json.Unmarshal([]byte(data), &jsonData)
	c.JSON(http.StatusOK, jsonData)
}

func (h *DataHandler) PostHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	data := map[string]any{}
	_ = json.Unmarshal(body, &data)

	key := fmt.Sprintf("%v", data["key"])
	err := h.service.SetServiceData(c.Request.Context(), key, string(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
