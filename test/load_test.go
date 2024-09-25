package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestHighLoad(t *testing.T) {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 2 * time.Second
	targetURL := fmt.Sprintf("http://localhost:%s/db-order/%s", cfg.PORT, testOrder.OrderUID)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    targetURL,
	})

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "High Load Test") {
		metrics.Add(res)
	}
	metrics.Close()

	report := &bytes.Buffer{}
	vegeta.NewTextReporter(&metrics).Report(report)
	fmt.Println(report.String())

	if metrics.Success < 0.95 {
		t.Errorf("Less than 95%% of requests were successful: %.2f%%", metrics.Success*100)
	}

	if metrics.Latencies.P95 > 200*time.Millisecond {
		t.Errorf("95th percentile latency is too high: %s", metrics.Latencies.P95)
	}
}

func TestParallelRequests(t *testing.T) {
	rate := vegeta.Rate{Freq: 50, Per: time.Second}
	duration := 5 * time.Second
	targetURL := fmt.Sprintf("http://localhost:%s/db-order/%s", cfg.PORT, testOrder.OrderUID)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    targetURL,
	})

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Parallel Requests Test") {
		metrics.Add(res)
	}
	metrics.Close()

	report := &bytes.Buffer{}
	vegeta.NewTextReporter(&metrics).Report(report)
	fmt.Println(report.String())

	if metrics.Success < 0.90 {
		t.Errorf("Less than 90%% of requests were successful: %.2f%%", metrics.Success*100)
	}

	if metrics.Latencies.P95 > 300*time.Millisecond {
		t.Errorf("95th percentile latency is too high: %s", metrics.Latencies.P95)
	}
}

func TestServiceAvailability(t *testing.T) {
	targetURL := fmt.Sprintf("http://localhost:%s/db-order/%s", cfg.PORT, testOrder.OrderUID)
	response, err := http.Get(targetURL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}
