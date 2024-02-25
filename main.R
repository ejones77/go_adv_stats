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
  elapsed <- Sys.time() - start
  memUsage <- mem_used()
  
  result <- list(
    language = "R",
    num_samples = numSamples,
    processing_time = as.character(elapsed),
    memory_usage = as.numeric(memUsage)
  )
  
  results[[length(results) + 1]] <- result
  
  print(paste("Mean of bootstrap samples:", mean(bootResult$t)))
  print(paste("Standard deviation of bootstrap samples:", sd(bootResult$t)))
}

write_json(results, "results_R.json")