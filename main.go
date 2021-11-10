package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/krishh1at/config"
	"github.com/krishh1at/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	err = config.DbConfig()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("MongoDB API")
	log.Println("Server is getting started...")

	r := routers.Router()
	log.Println("Server has started at port: 8080. To exit please press Ctr+C.")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
