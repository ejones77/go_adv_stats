package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Result struct {
	Language       string  `json:"language"`
	NumSamples     int     `json:"num_samples"`
	ProcessingTime float64 `json:"processing_time"`
	MemoryUsage    uint64  `json:"memory_usage"`
}

func main() {
	rResults := readResults("evaluate/results_R.json")
	goResults := readResults("evaluate/results_go.json")

	for i, rResult := range rResults {
		goResult := goResults[i]

		fmt.Printf("For %d samples:\n", rResult.NumSamples)

		timeRatio := rResult.ProcessingTime / goResult.ProcessingTime
		if rResult.ProcessingTime < goResult.ProcessingTime {
			fmt.Printf("R is %.2f times faster than Go.\n", timeRatio)
		} else if rResult.ProcessingTime > goResult.ProcessingTime {
			fmt.Printf("Go is %.2f times faster than R.\n", timeRatio)
		} else {
			fmt.Println("R and Go had the same processing time.")
		}

		memRatio := float64(rResult.MemoryUsage) / float64(goResult.MemoryUsage)
		if rResult.MemoryUsage < goResult.MemoryUsage {
			fmt.Printf("R uses %.2f times less memory than Go.\n", memRatio)
		} else if rResult.MemoryUsage > goResult.MemoryUsage {
			fmt.Printf("Go uses %.2f times less memory than R.\n", memRatio)
		} else {
			fmt.Println("R and Go used the same amount of memory.")
		}

		fmt.Println()
	}
}

func readResults(filename string) []Result {
	file, _ := os.ReadFile(filename)

	var results []Result
	json.Unmarshal(file, &results)

	for _, result := range results {
		fmt.Printf("Result: %+v\n", result)
	}

	return results
}
