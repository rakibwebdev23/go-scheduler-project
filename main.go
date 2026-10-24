package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID          string
	Name        string
	IsRecurring bool
	stopChan    chan struct{}
	wg          sync.WaitGroup
}

type Scheduler struct {
	tasks    map[string]*Task
	mu       sync.Mutex
	globalWg sync.WaitGroup
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

func generateTaskID() string {
	return fmt.Sprintf("task-%d", time.Now().UnixNano())
}

func (s *Scheduler) ScheduleOnce (name string, delay time.Duration, action func()) string{
	taskID := generateTaskID()

	fmt.Printf("[%s] SCHEDULER: Scheduling one-off task '%s' (ID: %s) to run after %s\n", time.Now().Format("15:04:05.000"), name, taskID, delay)

	s.globalWg.Add(1)
	time.AfterFunc(delay, func () {
		defer s.globalWg.Done()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Exceuting one of action.\n", 
		time.Now().Format("15:04:05.000"), name, taskID)

		action()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Finished one of action.\n", 
		time.Now().Format("15:04:05.000"), name, taskID)
	})
	return  taskID
}

func main() {

}
