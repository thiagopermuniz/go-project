package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projeto/internal/service"
)

type DataResponse struct {
	Data string `json:"data"`
}

type DataHandler struct {
	service service.DataServiceInterface
}

func NewDataHandler(svc service.DataServiceInterface) *DataHandler {
	return &DataHandler{
		service: svc,
	}
}

func (h *DataHandler) GetHello(c *gin.Context) {
	key := c.Query("key")
	data, err := h.service.GetServiceData(c.Request.Context(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get data"})
		fmt.Println(err)
		return
	}

	response := DataResponse{Data: data}
	c.JSON(http.StatusOK, response)
}

func (h *DataHandler) SetHello(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	data, err := h.service.SetServiceData(c.Request.Context(), key, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get data"})
		fmt.Println(err)
		return
	}

	response := DataResponse{Data: data}
	c.JSON(http.StatusOK, response)
}
