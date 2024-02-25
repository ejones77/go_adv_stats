package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/floats"
)

func TestBootstrap(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	numSamples := 100

	samples := bootstrap(data, numSamples)

	assert.Equal(t, numSamples, len(samples), "Expected number of samples to be equal to numSamples")

	min, max := floats.Min(data), floats.Max(data)
	for _, sample := range samples {
		assert.True(t, min <= sample && sample <= max, "Expected sample to be within range of original data")
	}
}
