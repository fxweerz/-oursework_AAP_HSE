package main

import (
	"log"

	db "task-prioritizer/internal/database"
	router "task-prioritizer/internal/transport"
)

func main() {

	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter()

	r.Run(":8080")
}