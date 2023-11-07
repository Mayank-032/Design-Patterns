package main

import (
	"fmt"
	"singleton-pattern/config"
	"sync"
	"testing"
)

func TestMainWithMultipleThreads(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			cfg := config.GetConfig()
			fmt.Printf("Goroutine %d: Environment: %s, Port: %s\n", id, cfg.Environment, cfg.Port)
		}(i)
	}

	wg.Wait()
}