# go_adv_stats

Submitted assignment for Northwestern MSDS 431


This repository compares the performance of Go and R in bootstrap sampling. The program generates bootstrap samples from a given dataset, calculates the mean of each sample, and records the processing time and memory usage for both languages.

## Overview

`main.go` and `main.R` are set up to perform the same series of tasks 
- perform bootstrap sampling on the same array for 1000, 100,000, and 1,000,000 samples
- record processing time and memory usage, storing results in a JSON file in the `evaluate` subdirectory
- prints the sample mean and sample standard deviation to the terminal

The bootstrap algorithm in R is straightforward, the `boot` library has a function that works right out of the box to generate sample statistics. 

It took a little more work upfront to code the algorithm in Go. 

The `bootstrap` function in `main.go` takes in a `data`, a slice of floats and `numSamples` for the number of samples to generate. To generate a bootstrap sample, it runs through two nested for loops:

- The outer loop runs `numSamples` times. Each iteration generates a bootstrap sample and calculates its mean.

- The inner loop runs equal to the size of `data`. Each iteration samples the original data with replacement to get a new temporary set of numbers to form the bootstrap statistic, which is the mean of what the inner loop generates.

`evaluate/evaluate.go` compares the two JSON files and prints results to the terminal
- calculates how many times faster or more efficient one language is over the other.
- compares cloud costs based on the hourly compute rate, for this example I'm using the `t4g.medium` EC2 instance [from AWS](https://aws.amazon.com/ec2/pricing/on-demand/)
- The cost calculation comes from taking the execution time and multiplying by 500,000 to estimate average execution time at scale, then obtaining the cost from the hourly rate.



## Dependencies

For Go
- gonum
- encoding/json
- runtime

For R
- boot
- pryr
- jsonlite


## Usage
Ensure R dependencies are installed, this command should do the trick

```Rscript -e "install.packages(c('boot', 'jsonlite', 'pryr'), repos='https://cran.rstudio.com/')"```

Run the Go and R programs to generate results 

```go run main.go```

```Rscript run main.R```


Then execute the evaluate program to confirm results
```go run .\evaluate\evaluate.go```

## Results

After performing the test on my local machine, these were my results:

```
For 1000 samples:
Go is +Inf times faster than R.
Go uses 152.82 times less memory than R.
R cost: $0.0560
Go cost: $0.0000
Cost savings with Go: $0.0560

For 100000 samples:
Go is 38.12 times faster than R.
Go uses 12.46 times less memory than R.
R cost: $2.8467
Go cost: $0.0747
Cost savings with Go: $2.7720

For 1000000 samples:
Go is 50.23 times faster than R.
Go uses 4.03 times less memory than R.
R cost: $31.8780
Go cost: $0.6347
Cost savings with Go: $31.2433
```

The "infinitely faster" result for 1000 samples was unexpected but checks out because the execution time was less than 1 millisecond in my case.

Go maintains an impressive performance boost over R in both speed and memory usage. Go's performance boost strengthens as the scale grows, but Go is less of an improvement in memory usage at scale. Still, it's impressive to be 4x more efficient with 1 million samples. 

Because Go's performance boost grows exponentially, so too do the cloud savings costs. You'll spend $30 with R before you spend a single dollar with Go. The bigger the data you need to process, the more appealing Go is as an alterative.  