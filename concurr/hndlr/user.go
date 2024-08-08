package handler

import (
	"github.com/gin-gonic/gin"
	"main/httpcli"
)

func UserHandler(client httpcli.CustomClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c
		client.Get(c.Request.Context())
	}

}
