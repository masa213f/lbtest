package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.999Z"

var (
	interval  = 10 * time.Millisecond
	timeout   = 10 * time.Millisecond
	targetURL string
)

func main() {
	if str := os.Getenv("INTERVAL"); str != "" {
		d, err := time.ParseDuration(str)
		if err != nil {
			log.Fatalf("invalid INTERVAL: %v", err)
		}
		interval = d
	}
	if str := os.Getenv("TIMEOUT"); str != "" {
		d, err := time.ParseDuration(str)
		if err != nil {
			log.Fatalf("invalid TIMEOUT: %v", err)
		}
		timeout = d
	}
	targetURL := os.Getenv("TARGET")
	if targetURL == "" {
		log.Fatal("no TARGET")
	}
	log.Printf("Hello! target %s, interval %s, timeout %s", targetURL, interval, timeout)
	log.SetFlags(0)

	c := &http.Client{Timeout: timeout}
	tick := time.NewTicker(interval)
	errorCount := 0
	for {
		start := time.Now()
		resp, err := c.Get(targetURL)
		d := time.Since(start)
		if err != nil {
			log.Printf("%s NG %v", start.Format(timeFormat), err)
			errorCount++
		} else {
			log.Printf("%s OK %v", start.Format(timeFormat), d)
			resp.Body.Close() // we don't use the response body.
		}

		if errorCount > 5000 {
			break
		}
		select {
		case <-tick.C:
		}
	}

	log.SetFlags(log.LstdFlags)
	log.Printf("Bye!")
}
