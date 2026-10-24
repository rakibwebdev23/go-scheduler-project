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

func main() {

}
