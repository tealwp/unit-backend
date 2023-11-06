package cfg

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int
	GRPC_PORT   int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	grpc_port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	return &Config{
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     db_port,
		GRPC_PORT:   grpc_port,
	}
}
