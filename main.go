package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/NessibeliY/mechta.kz/conf"
)

func main() {
	config, err := conf.New()
	if err != nil {
		log.Fatal(err)
	}

	numGoroutines := 1
	if len(os.Args) > 1 {
		numGoroutines, err = strconv.Atoi(os.Args[1])
		if err != nil || numGoroutines < 1 {
			numGoroutines = 1
		}
	}

	fmt.Printf("Number of goroutines: %d\n", numGoroutines)

	totalSum := parallelSum(config.NumberSets, numGoroutines)
	fmt.Printf("Total sum: %d\n", totalSum)
}

func parallelSum(numberSets []conf.NumberSet, numGoroutines int) int {
	batchSize := (len(numberSets) + numGoroutines - 1) / numGoroutines

	results := make(chan int, numGoroutines)

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		start := i * batchSize
		end := start + batchSize

		if end > len(numberSets) {
			end = len(numberSets)
		}

		go func(start, end int) {
			defer wg.Done()

			partialSum := 0
			for _, numberSet := range numberSets[start:end] {
				partialSum += numberSet.A + numberSet.B
			}

			results <- partialSum
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for partialSum := range results {
		totalSum += partialSum
	}

	return totalSum
}
