package workers

import "sync"

// Createa  worker pool 
func StartWorkerPool[T any](
	totalWorkers int,
	jobs []T,
	jobFunc func(T),
) {
	jobChan := make(chan T)

	var wg sync.WaitGroup

	// Create workers
	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)

		go worker(&wg, jobChan, jobFunc)
	}

	// Send jobs
	for _, j := range jobs { 
		jobChan <- j
	}

	close(jobChan)

	wg.Wait()
}

func worker[T any](
	wg *sync.WaitGroup,
	jobChan chan T,
	jobFunc func(T),
) {
	defer wg.Done()

	for job := range jobChan {
		jobFunc(job)
	}
}
