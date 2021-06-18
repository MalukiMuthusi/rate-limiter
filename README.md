# Bursty Rate Limiter

Go server that sends batches of numbers base of the given input.
Given the input:

```txt
skipBatches = 0
printBatches = 4
batchSize = 3
toAdd = 5
```

It will output:

```txt
[0,5,10]
[15,20,25]
[25,30,35],
[40,45,50]
```

To run the program:

- `go build -o rate-limiter .`
- then `./rate-limiter.o < input.txt`
