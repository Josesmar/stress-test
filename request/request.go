package request

import (
	"net/http"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
}

func PerformRequest(client *http.Client, url string) Result {
	startReq := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{StatusCode: 0, Duration: time.Since(startReq)}
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 ...")

	resp, err := client.Do(req)
	if err != nil {
		return Result{StatusCode: 0, Duration: time.Since(startReq)}
	}
	defer resp.Body.Close()

	return Result{StatusCode: resp.StatusCode, Duration: time.Since(startReq)}
}
