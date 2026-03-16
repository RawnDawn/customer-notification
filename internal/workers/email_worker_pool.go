package workers

import (
	"sync"

	"github.com/rawndawn/customer-notification/internal/email"
	"github.com/rawndawn/customer-notification/internal/models"
)

// Worker pool to send promotional email to the customers
func StartPromotionalEmailWorkerPool(
	totalWorkers int,
	customers []models.Customer,
){
	// Channel without buffer to block until a worker get ready to recieve another job
	jobs := make(chan *models.Customer)
	
	var wg sync.WaitGroup

	// Create workers
	for i := 0; i < totalWorkers; i ++ { 
		wg.Add(1)

		go worker(&wg, jobs)
	}

	// Send jobs
	for index := range customers { 
		jobs <- &customers[index]
	}

	close(jobs)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, jobs chan *models.Customer) {
	defer wg.Done()

	for customer := range jobs { 
		email.SendPromotional(customer)
	}
}