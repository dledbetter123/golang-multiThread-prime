package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	pathname = flag.String("pathname", "", "Path to the binary file")
	M        = flag.Int("M", 1, "Number of worker threads")
	N        = flag.Int("N", 64*1024, "Size of each segment in bytes")
	C        = flag.Int("C", 1024, "Chunk size in bytes")
)

type Job struct {
	Pathname string
	Start    int64
	Length   int
}

type Result struct {
	Job   Job
	Count int
}

func isPrime(n uint64) bool {
	if n <= 1 {
		return false
	}
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	randSleep := time.Millisecond * time.Duration(rand.Intn(200)+400)
	time.Sleep(randSleep)
	for job := range jobs {
		file, err := os.Open(job.Pathname)
		if err != nil {
			log.Fatalf("Worker %d: Failed to open file: %v", id, err)
		}
		defer file.Close()

		primesCount := 0
		buffer := make([]byte, *C)
		for i := 0; i < job.Length; i += *C {
			readSize := *C
			if job.Length-i < *C {
				readSize = job.Length - i
			}
			_, err := file.ReadAt(buffer[:readSize], job.Start+int64(i))
			if err != nil {
				log.Fatalf("Worker %d: Failed to read file: %v", id, err)
			}

			for j := 0; j < readSize; j += 8 {
				var num uint64
				binary.LittleEndian.PutUint64(buffer[j:j+8], num)
				if isPrime(num) {
					primesCount++
				}
			}
		}

		result := Result{Job: job, Count: primesCount}
		results <- result
		log.Printf("Worker %d completed job: %+v with result: %+v\n", id, job, result)
	}
}

func dispatcher(pathname string, fileSize int64, jobs chan<- Job) {
	for start := int64(0); start < fileSize; start += int64(*N) {
		length := *N
		if start+int64(length) > fileSize {
			length = int(fileSize - start)
		}
		job := Job{Pathname: pathname, Start: start, Length: length}
		jobs <- job
	}
	close(jobs)
}

func consolidator(results <-chan Result, done chan<- int) {
	totalPrimes := 0
	for result := range results {
		totalPrimes += result.Count
	}
	done <- totalPrimes
}

func main() {
	flag.Parse()

	if *pathname == "" {
		log.Fatal("Pathname is required")
	}

	fileInfo, err := os.Stat(*pathname)
	if err != nil {
		log.Fatalf("Failed to get file stats: %v", err)
	}

	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	done := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < *M; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go dispatcher(*pathname, fileInfo.Size(), jobs)
	go consolidator(results, done)

	totalPrimes := <-done
	fmt.Printf("Total primes: %d\n", totalPrimes)
}
