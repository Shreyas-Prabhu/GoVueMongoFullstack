package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("---welcome to Go server---")

	//Get environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port:= os.Getenv("PORT")
	movieRouter := router.NewRouter()
	movieRouter.Run(":" + port)
}
