library(boot)
library(jsonlite)
library(pryr)

statistic <- function(data, indices) {
  return(mean(data[indices]))
}

data <- c(1.2, 2.5, 3.7, 4.1, 5.3, 6.8, 7.4, 8.9, 9.2, 10.5)
sampleSizes <- c(1000, 100000, 1000000)
results <- list()

for (numSamples in sampleSizes) {
  start <- Sys.time()
  bootResult <- boot(data, statistic, R = numSamples)
  elapsed <- difftime(Sys.time(), start, units = "secs") * 1000
  memUsage <- mem_used()
  
  result <- list(
    language = unbox("R"),
    num_samples = unbox(numSamples),
    # Go rounds down when computing milliseconds, we should do the same
    processing_time = unbox(floor(as.numeric(elapsed))),
    memory_usage = unbox(as.numeric(memUsage))
  )
  
  results[[length(results) + 1]] <- result
  
  print(paste("Mean of bootstrap samples:", mean(bootResult$t)))
  print(paste("Standard deviation of bootstrap samples:", sd(bootResult$t)))
}

write_json(results, "evaluate/results_R.json")