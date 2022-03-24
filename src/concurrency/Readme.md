![img_1.png](img_1.png)
##This demonstrates one of the ways where you can use concurrency in golang
1. make a slow fib function
2. make worker threads to process it

you should try:
1. playing with the number of workers
2. playing with number of jobs
3. create a deadlock when workers and jobs are not the same, then try to understand why (a: its all about the buffer channel space!)
4. add more go workers, reduce go workers