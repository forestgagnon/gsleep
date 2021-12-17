package main

import (
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("duration in Go duration syntax is required")
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		log.Fatalf("duration in Go duration syntax is required: %s", err.Error())
	}
	time.Sleep(duration)
}
