package main

import "fmt"

// preface
// channels are created by the use of <- and there is no such thing as "->" so its easy to remember
// if the arrow is in front of the var like i<- then i channel is recieving a signal
// if you are doing a <- i, then you are sending a signal to a channel from i channel

func main() {
	jobs := make(chan int, 100)    // create a signal channel with a buffer space of 100 with intent to create jobs, ill be sending ints through this channel
	results := make(chan int, 100) // create a signal channel with a buffer space of 100 with intend to recieve results in int, ill be sending ints through this channel

	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)
	go worker(jobs, results) // non blocking line, because its a goroutine (kinda like creating a thread and going .run() in java)

	for i := 0; i < 100; i++ {
		jobs <- i // i == used in fib(i), creates a job here (i) and sends that as a signal to the worker through the jobs channel (the workers one lines 9-15)
	}
	close(jobs) // blocking call, eventually unblocks after feeding all the jobs (for loop), so since now im not going to give it anymore jobs so close the channel
	// that means worker doesnt expect to recieve anymore and the worker will terminate because it notices the jobs channel is closed, see func worker

	for i := 0; i < 100; i++ {
		fmt.Println(<-results) // keeps recieving results int until all the jobs are done
	}

}
func worker(jobs <-chan int, results chan<- int) { // workers func params: jobs <- expects an inbound jobs channel (a recv from channel), and results chan<-: outbound results channel (send to channel). this is good practice to ensure you are sending the correct type of channel to this function, otherwise the arrows are optional

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
