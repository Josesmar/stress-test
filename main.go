package main

import (
	"stress-test/config"
	"stress-test/report"
	"stress-test/request"
	"sync"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	client := request.NewClient()

	results := make(chan request.Result, cfg.TotalRequests)
	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < cfg.TotalRequests/cfg.Concurrency; j++ {
				results <- request.PerformRequest(client, cfg.URL)
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	report.GenerateReport(results)
}
