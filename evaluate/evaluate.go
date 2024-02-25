package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Result struct {
	Language       string `json:"language"`
	NumSamples     int    `json:"num_samples"`
	ProcessingTime string `json:"processing_time"`
	MemoryUsage    uint64 `json:"memory_usage"`
}

func main() {
	rResults := readResults("results_R.json")
	goResults := readResults("results_Go.json")

	for i, rResult := range rResults {
		goResult := goResults[i]

		rDuration, _ := time.ParseDuration(rResult.ProcessingTime)
		goDuration, _ := time.ParseDuration(goResult.ProcessingTime)

		fmt.Printf("For %d samples:\n", rResult.NumSamples)

		if rDuration < goDuration {
			fmt.Println("R was faster.")
		} else {
			fmt.Println("Go was faster.")
		}

		if rResult.MemoryUsage < goResult.MemoryUsage {
			fmt.Println("R used less memory.")
		} else {
			fmt.Println("Go used less memory.")
		}

		fmt.Println()
	}
}

func readResults(filename string) []Result {
	file, _ := os.ReadFile(filename)

	var results []Result
	json.Unmarshal(file, &results)

	return results
}
