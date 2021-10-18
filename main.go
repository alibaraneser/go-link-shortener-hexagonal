package main

import (
	"github.com/alibaraneser/go-link-shortener-hexagonal/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}