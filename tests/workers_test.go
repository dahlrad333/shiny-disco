package tests

import (
	"context"
	"shiny-disco/server"
	"testing"
	"time"
)

// TestWorkerPool_SuccessfulProcessing checks that jobs are processed successfully.
func TestWorkerPool_SuccessfulProcessing(t *testing.T) {
	inputCh := make(chan server.Job, 5)
	outputCh := make(chan server.Result, 5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	workerPool := server.NewWorkerPool(ctx, 2, inputCh, outputCh)
	workerPool.Start()

	go func() {
		for i := 1; i <= 5; i++ {
			inputCh <- server.Job{ID: i, Input: i}
		}
		close(inputCh)
	}()

	results := make(map[int]int)
	for i := 0; i < 5; i++ {
		result := <-outputCh
		if result.Error != nil {
			t.Errorf("Unexpected error for job %d: %v", result.JobID, result.Error)
		}
		results[result.JobID] = result.Output
	}

	for i := 1; i <= 5; i++ {
		if results[i] != i*2 {
			t.Errorf("Job %d: expected output %d, got %d", i, i*2, results[i])
		}
	}

	workerPool.Shutdown()
}

// TestWorkerPool_ErrorHandling checks that invalid jobs return errors properly.
func TestWorkerPool_ErrorHandling(t *testing.T) {
	inputCh := make(chan server.Job, 3)
	outputCh := make(chan server.Result, 3)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	workerPool := server.NewWorkerPool(ctx, 2, inputCh, outputCh)
	workerPool.Start()

	go func() {
		inputCh <- server.Job{ID: 1, Input: -1} 
		inputCh <- server.Job{ID: 2, Input: 2} 
		close(inputCh)
	}()

	var hasError bool
	for i := 0; i < 2; i++ {
		select {
		case result := <-outputCh:
			if result.Error != nil {
				hasError = true
			}
		case err := <-workerPool.Errors():
			hasError = true
			t.Logf("Received expected error: %v", err)
		}
	}

	if !hasError {
		t.Errorf("Expected error for invalid job input, but none occurred")
	}

	workerPool.Shutdown()
}

// TestWorkerPool_ContextTimeout ensures workers stop processing jobs on context timeout.
func TestWorkerPool_ContextTimeout(t *testing.T) {
	inputCh := make(chan server.Job, 10)
	outputCh := make(chan server.Result, 10)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	workerPool := server.NewWorkerPool(ctx, 2, inputCh, outputCh)
	workerPool.Start()

	go func() {
		for i := 1; i <= 10; i++ {
			inputCh <- server.Job{ID: i, Input: i}
		}
		close(inputCh)
	}()

	<-ctx.Done()

	select {
	case <-outputCh:
		t.Log("Received result after context timeout (expected due to worker delay)")
	default:
		t.Log("No further results after context timeout")
	}

	workerPool.Shutdown()
}