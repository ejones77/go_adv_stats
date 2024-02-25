package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"gonum.org/v1/gonum/stat"
)

type Result struct {
	Language       string  `json:"language"`
	NumSamples     int     `json:"num_samples"`
	ProcessingTime float64 `json:"processing_time"`
	MemoryUsage    uint64  `json:"memory_usage"`
}

func bootstrap(data []float64, numSamples int) []float64 {
	sampleSize := len(data)
	samples := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		sample := make([]float64, sampleSize)
		for j := 0; j < sampleSize; j++ {
			index := rand.Intn(sampleSize)
			sample[j] = data[index]
		}
		samples[i] = stat.Mean(sample, nil)
	}

	return samples
}

func main() {
	data := []float64{1.2, 2.5, 3.7, 4.1, 5.3, 6.8, 7.4, 8.9, 9.2, 10.5}
	sampleSizes := []int{1000, 100000, 1000000}
	results := []Result{}

	for _, numSamples := range sampleSizes {
		start := time.Now()
		samples := bootstrap(data, numSamples)
		elapsed := time.Since(start)

		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		result := Result{
			Language:       "Go",
			NumSamples:     numSamples,
			ProcessingTime: float64(elapsed.Milliseconds()),
			MemoryUsage:    memStats.Alloc,
		}

		results = append(results, result)

		mean := stat.Mean(samples, nil)
		stdev := stat.StdDev(samples, nil)

		fmt.Printf("For %d samples, Mean: %v, StdDev: %v\n", numSamples, mean, stdev)
	}

	file, _ := os.Create("evaluate/results_go.json")
	defer file.Close()

	json.NewEncoder(file).Encode(results)
}
