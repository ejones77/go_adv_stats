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

	// using AWS EC2 t4g.medium as an example
	hourlyRate := 0.0336

	// scaled up threshold of how many times the program is run
	numRuns := 500000

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

		// Converting ms to hours
		rCost := (rResult.ProcessingTime / 3600000) * hourlyRate
		goCost := (goResult.ProcessingTime / 3600000) * hourlyRate

		fmt.Printf("R cost: $%.4f\n", rCost*float64(numRuns))
		fmt.Printf("Go cost: $%.4f\n", goCost*float64(numRuns))

		costSavings := (rCost - goCost) * float64(numRuns)
		fmt.Printf("Cost savings with Go: $%.4f\n", costSavings)

		fmt.Println()
	}
}

func readResults(filename string) []Result {
	file, _ := os.ReadFile(filename)

	var results []Result
	json.Unmarshal(file, &results)

	return results
}
