package workers

import (
	"log"
	"sync"

	"github.com/rawndawn/customer-notification/internal/models"
)

func StartPromotionalEmailWorkerPool(
	totalWorkers int,
	customers []models.Customer,
	job func(models.Customer),
){
	// Channel without buffer to block until a worker get ready to recieve another job
	jobs := make(chan models.Customer)
	
	var wg sync.WaitGroup

	// Create workers
	for i := 0; i < totalWorkers; i ++ { 
		wg.Add(1)

		go sendEmail(&wg, jobs)
	}

	// Send jobs
	for _, customer := range customers { 
		jobs <- customer
	}

	close(jobs)

	wg.Wait()
}

func sendEmail(wg *sync.WaitGroup, jobs chan models.Customer) {
	defer wg.Done()

	for customer := range jobs { 
		log.Println("Sending email...", customer.Email)
	}

}