package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"projeto/internal/handlers"
	"projeto/internal/repository"
	"projeto/internal/service"
)

func main() {
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")
	repo, err := repository.NewRedisRepository(redisEndpoint)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	svc := service.NewDataService(repo)
	dataHandler := handlers.NewDataHandler(svc)

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/hello", dataHandler.GetHello)
	r.GET("/hellow", dataHandler.SetHello)

	svrErr := r.Run(":8080")
	if svrErr != nil {
		panic(svrErr)
	}
}
