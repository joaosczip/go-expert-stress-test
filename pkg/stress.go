package stress

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type StressTester struct {
	concurrency int
	requests    int
	url         string
	timeout     time.Duration
}

type response struct {
	statusCode int
}

func NewStressTester(concurrency, requests int, url string, timeout time.Duration) *StressTester {
	return &StressTester{
		concurrency: concurrency,
		requests:    requests,
		url:         url,
		timeout:     timeout,
	}
}

func (s *StressTester) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)

	defer cancel()

	guard := make(chan struct{}, s.concurrency)
	responses := make(chan response)
	errorsCh := make(chan error)

	for i := 0; i < s.requests; i++ {
		guard <- struct{}{}
		go s.doRequest(ctx, i, guard, responses, errorsCh)
	}

	close(guard)

	statusCodes := make(map[int]int)
	successfullRequests := 0

	for i := 0; i < s.requests; i++ {
		select {
		case response := <-responses:
			statusCodes[response.statusCode]++
			if response.statusCode >= 200 && response.statusCode < 300 {
				successfullRequests++
			}
		case err := <-errorsCh:
			if errors.Is(err, context.DeadlineExceeded) {
				statusCodes[504]++
			} else {
				statusCodes[500]++
			}
		}
	}

	close(responses)
	close(errorsCh)

	for status, total := range statusCodes {
		fmt.Printf("Status %d: %d\n", status, total)
	}

	fmt.Printf("Total requests: %d\n", s.requests)
	fmt.Printf("Successfull requests: %d\n", successfullRequests)
	fmt.Printf("Failed requests: %d\n", s.requests-successfullRequests)
}

func (s *StressTester) doRequest(ctx context.Context, i int, guard <-chan struct{}, responses chan<- response, errorsCh chan<- error) {
	fmt.Printf("Processing request %d\n", i)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.url, nil)

	if err != nil {
		<-guard
		errorsCh <- err
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		<-guard
		errorsCh <- err
		return
	}

	<-guard
	responses <- response{statusCode: resp.StatusCode}
}
