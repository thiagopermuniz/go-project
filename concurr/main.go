package main

import "github.com/gin-gonic/gin"

func main() {
	svr := gin.Default()
	svr.GET("/", handler)

}
