package main

import (
	"log"
	"subscriptions-service/internal/config"
	"subscriptions-service/internal/handler"
	"subscriptions-service/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatal("DB connect error:", err)
	}
	defer db.Close()
	log.Println("DB connected successfully")

	h := handler.NewHandler(repository.NewSubscriptionRepo(db))

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	h.RegisterRoutes(r)

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
