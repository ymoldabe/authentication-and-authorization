package main

import (
	"auth_in_go/models"
	"auth_in_go/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if err := models.InitDB(config); err != nil {
		log.Fatal(err)
	}

	routers.AuthRouters(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
