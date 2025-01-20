package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	SendGridAPIKey    string
	SendGridFromEmail string
}

func initConfig() Config {
	return Config{
		SendGridAPIKey:    getEnvOrPanic("SENDGRID_API_KEY", "SendGrid API KEY is required"),
		SendGridFromEmail: getEnvOrPanic("SENDGRID_FROM_EMAIL", "SendGrid From email is required"),
	}
}

func getEnvOrPanic(key, err string) string {
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(err)
}
