package server

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Input int
}

type Result struct {
	JobID  int
	Output int
	Error  error
}

type WorkerPool struct {
	ctx       context.Context
	cancel    context.CancelFunc
	inputCh   chan Job
	outputCh  chan Result
	errorCh   chan error
	workerNum int
	wg        sync.WaitGroup
}

func NewWorkerPool(ctx context.Context, workerNum int, inputCh chan Job, outputCh chan Result) *WorkerPool {
	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		ctx:       ctx,
		cancel:    cancel,
		inputCh:   inputCh,
		outputCh:  outputCh,
		errorCh:   make(chan error, workerNum),
		workerNum: workerNum,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerNum; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id))) 

	for {
		select {
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d: shutting down\n", id)
			return
		case job, ok := <-wp.inputCh:
			if !ok {
				fmt.Printf("Worker %d: input channel closed\n", id)
				return
			}
			
			processingTime := time.Duration(rng.Intn(5)) * time.Second
			fmt.Printf("Worker %d: processing job %d, will take %v\n", id, job.ID, processingTime)
			time.Sleep(processingTime)

			result, err := wp.processJob(job)
			if err != nil {
				wp.errorCh <- fmt.Errorf("Worker %d: job %d failed: %v", id, job.ID, err)
				continue
			}
			
			wp.outputCh <- result
			fmt.Printf("Worker %d: completed job %d\n", id, job.ID)
		}
	}
}

func (wp *WorkerPool) processJob(job Job) (Result, error) {
	if job.Input < 0 {
		return Result{}, fmt.Errorf("invalid input: %d", job.Input)
	}
	
	return Result{JobID: job.ID, Output: job.Input * 2}, nil
}

func (wp *WorkerPool) Shutdown() {
	wp.cancel() 
	wp.wg.Wait() 
	close(wp.errorCh)
	close(wp.outputCh)
	fmt.Println("Worker pool shut down gracefully")
}

func (wp *WorkerPool) Errors() <-chan error {
	return wp.errorCh
}