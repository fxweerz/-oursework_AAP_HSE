package main

import (
	"task-prioritizer/internal/transport"

	"task-prioritizer/internal/database"
)

func main() {
	db.Init()
	r := transport.SetupRouter()
	r.Run(":8080")
}
