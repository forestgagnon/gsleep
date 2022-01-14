package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("duration in Go duration syntax or number of seconds is required")
	}

	time.Sleep(duration(os.Args[1]))
}

func duration(input string) time.Duration {
	if seconds, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
		return time.Duration(seconds) * time.Second
	}

	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		log.Fatalf("invalid Go duration syntax: %s", err.Error())
	}

	return duration
}
