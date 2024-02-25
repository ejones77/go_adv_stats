# go_adv_stats

Submitted assignment for Northwestern MSDS 431


This repository compares the performance of Go and R in bootstrap sampling. The program generates bootstrap samples from a given dataset, calculates the mean of each sample, and records the processing time and memory usage for both languages.

## Overview

`main.go` and `main.R` are set up to perform the same series of tasks 
- perform bootstrap sampling on the same array for 1000, 100,000, and 1,000,000 samples
- record processing time and memory usage, storing results in a JSON file in the `evaluate` subdirectory
- prints the sample mean and sample standard deviation to the terminal

`evaluate/evaluate.go` compares the two JSON files and prints results to the terminal, calculating how many times faster or more efficient one language is over the other.

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

For 100000 samples:
Go is 38.12 times faster than R.
Go uses 12.46 times less memory than R.

For 1000000 samples:
Go is 50.23 times faster than R.
Go uses 4.03 times less memory than R.
```

The "infinitely faster" result for 1000 samples was unexpected but checks out because the execution time was less than 1 millisecond in my case.

Go maintains an impressive performance boost over R in both speed and memory usage. Go's performance boost strengthens as the scale grows, but Go is less of an improvement in memory usage at scale. Still, it's impressive to be 4x more efficient with 1 million samples. 
