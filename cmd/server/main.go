package main

import (
	"log"
	"mpango-wa-cuddles/internal/router"
)

func main() {
	r := router.New()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}


