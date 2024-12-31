package report

import (
	"fmt"
	"stress-test/request"
	"time"
)

func GenerateReport(results chan request.Result) {
	report := make(map[int]int)
	totalDuration := time.Duration(0)
	count := 0

	startTime := time.Now()

	for result := range results {
		report[result.StatusCode]++
		totalDuration += result.Duration
		count++
	}

	endTime := time.Now()

	testDuration := endTime.Sub(startTime)

	fmt.Println("End of test:", endTime)

	fmt.Printf("\nTest completed in %.2fs\n", testDuration.Seconds())
	fmt.Printf("Total requests: %d\n", count)

	avgDuration := testDuration.Seconds() / float64(count)
	fmt.Printf("Average time per request: %.6fs\n", avgDuration)

	for status, count := range report {
		fmt.Printf("Status %d: %d\n", status, count)
	}
}
