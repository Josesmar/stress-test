package config

import (
	"flag"
	"fmt"
	"time"
)

type Config struct {
	URL           string
	TotalRequests int
	Concurrency   int
}

func LoadConfig() Config {
	url := flag.String("url", "", "URL of the service to be tested")
	totalRequests := flag.Int("requests", 0, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of simultaneous calls")
	flag.Parse()

	if *url == "" || *totalRequests <= 0 || *concurrency <= 0 {
		panic("Invalid parameters. Use --url, --requests e --concurrency.")
	}

	start := time.Now()
	fmt.Println("Test start:", start)
	fmt.Printf("Starting load testing: URL=%s, Total Requests=%d, Competition=%d\n", *url, *totalRequests, *concurrency)

	return Config{
		URL:           *url,
		TotalRequests: *totalRequests,
		Concurrency:   *concurrency,
	}
}
