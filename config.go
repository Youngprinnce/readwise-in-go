package main

// import (
// 	"fmt"
// 	"os"
// )

// var Envs = initConfig()

// type Config struct {
// 	SendGridAPIKey    string
// 	SendGridFromEmail string
// }

// func initConfig() Config {
// 	return Config{
// 		SendGridAPIKey:    getEnvOrPanic("SENDGRID_API_KEY", "SendGrid API KEY is required"),
// 		SendGridFromEmail: getEnvOrPanic("SENDGRID_FROM_EMAIL", "SendGrid From email is required"),
// 	}
// }

// func getEnvOrPanic(key, err string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		fmt.Println("SENDGRID_API_KEY: ", value)
// 		return value
// 	}

// 	panic(err)
// }
