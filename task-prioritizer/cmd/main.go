package main

import (
	"log"

	"github.com/joho/godotenv"

	db "task-prioritizer/internal/database"
	router "task-prioritizer/internal/transport"
)

func main() {

	godotenv.Load()

	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter()

	r.Run(":8080")
}