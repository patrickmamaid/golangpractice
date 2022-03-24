package main

import "fmt"

const NUM_WORKERS int = 5
const NUM_JOBS int = 100

func main() {
	jobs := make(chan int, NUM_JOBS)
	results := make(chan int, NUM_JOBS)

	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine
	go worker(jobs, results) // non blocking line, because its a goroutine

	for i := 0; i < NUM_JOBS; i++ {
		jobs <- i // i == used in fib(i), creates a job here
	}
	close(jobs) // blocking call, eventually after feeding all the jobs, im not going to give it more jobs so close the channel
	// that means worker doesnt expect to recieve anymore

	for i := 0; i < NUM_JOBS; i++ {
		fmt.Println(<-results) // keeps recieving results int until all the jobs are done
	}

}
func worker(jobs <-chan int, results chan<- int) { // workers func expect inbound jobs channel (recv from channel), and outbound results channel (send to channel)

	for n := range jobs { // waits until it recieves a job, (how many jobs) goes into n
		results <- fib(n) // execute slow algo per n jobs, when done, send int through results channel
	}
}

func fib(n int) int { // really slow fib algo
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)

}
