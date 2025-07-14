package service

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Err        error
}

type Report struct {
	TotalRequests int
	Successful    int
	Errors        int
	StatusDist    map[int]int
	Duration      time.Duration
}

func RunStressTest(url string, requests, concurrency int, print func(a ...interface{})) {
	start := time.Now()
	results := make(chan Result, requests)
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < requests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := client.Get(url)
			if err != nil {
				results <- Result{StatusCode: 0, Err: err}
				return
			}
			defer resp.Body.Close()
			results <- Result{StatusCode: resp.StatusCode}
		}()
	}

	wg.Wait()
	close(results)

	report := processResults(results, time.Since(start))
	printReport(report, print)
}

func processResults(results <-chan Result, duration time.Duration) Report {
	total := 0
	successful := 0
	errors := 0
	statusDist := make(map[int]int)

	for result := range results {
		total++
		if result.Err != nil {
			errors++
		} else {
			successful++
			statusDist[result.StatusCode]++
		}
	}

	return Report{
		TotalRequests: total,
		Successful:    successful,
		Errors:        errors,
		StatusDist:    statusDist,
		Duration:      duration,
	}
}

func printReport(report Report, print func(a ...interface{})) {
	print("======================================")
	print("         RELATÓRIO DE STRESS TEST     ")
	print("======================================")
	print(fmt.Sprintf("Tempo total:         %v", report.Duration))
	print(fmt.Sprintf("Total de requests:   %d", report.TotalRequests))
	print(fmt.Sprintf("Sucesso (HTTP 2xx):  %d", report.Successful))
	print(fmt.Sprintf("Erros:               %d", report.Errors))
	print("--------------------------------------")
	print("Distribuição de status HTTP:")
	print(" Código | Quantidade")
	print("--------+-----------")

	codes := make([]int, 0, len(report.StatusDist))
	for code := range report.StatusDist {
		codes = append(codes, code)
	}
	sort.Ints(codes)

	for _, code := range codes {
		print(fmt.Sprintf(" %6d | %9d", code, report.StatusDist[code]))
	}
	print("======================================")
}
