package main

import "time"

/*
 * Complete the 'BurstyRateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 *  3. int batchSize
 *  4. int toAdd
 */

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {

	for j := 0; ; j += toAdd {
		<-requestChan
		// print the numbers in one batch
		i := j
		for k := 0; k < batchSize; k++ {
			time.Sleep(10 * time.Millisecond)
			resultChan <- i
			i += toAdd
		}

	}

}
