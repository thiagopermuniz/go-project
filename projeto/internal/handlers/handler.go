package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"projeto/internal/service"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var jsonData Data
	_ = json.Unmarshal([]byte(data), &jsonData)
	c.JSON(http.StatusOK, jsonData)
}

func (h *DataHandler) PostHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	data := Data{}
	_ = json.Unmarshal(body, &data)

	key := data.Key
	value := data.Value
	err := h.service.SetServiceData(c.Request.Context(), key, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := Data{Key: key, Value: value}
	c.JSON(http.StatusOK, response)
}
