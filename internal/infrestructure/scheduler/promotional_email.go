package scheduler

import "time"

type PromotionalEmailScheduler struct {
	hour   int
	minute int
	job    func()
}

// Create a Promotional Email Scheduler instance.
// Params:
// hour int - the hour that system will send the email
// minute int - the minute that system will send the email
// job func() - task to execute
func NewPromotionalEmailScheduler(hour, minute int, job func()) *PromotionalEmailScheduler {
	return &PromotionalEmailScheduler{
		hour:   hour,
		minute: minute,
		job:    job,
	}
}

// Start the promotional email scheduler. This sends an email every month
func (s *PromotionalEmailScheduler) Start() {
	// Endlessly for
	for {
		// Prepare first execution
		now := time.Now()
		nextRun := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			s.hour,
			s.minute,
			0,
			0,
			now.Location(),
		)

		// Executes every month
		if now.After(nextRun) {
			nextRun = nextRun.AddDate(0, 1, 0)
		}

		// Sleep until next run
		waitDuration := nextRun.Sub(now)
		time.Sleep(waitDuration)

		// Execute in a routine
		go s.job()
	}
}
