package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var program_start_time time.Time
var program_end_time time.Time
var total_completed_jobs = 0
var total_worker_busyness = 0

func main() {
	/*
		Goroutine interacts with other goroutines using channels
	*/
	var wg sync.WaitGroup
	//waitGroup pauses the termination of main prog so that
	//goroutines can perform the task and return back to main thread

	/*
			Three steps:
			1> Add : Specify how many go routines are there so that it waits until
			all the go routines are executed

			2> Wait: specified in the caller function when you want to wait for all the
			goroutines to complete the execution

			3> Done: specified in the called function and it means that this go routine
			has completed its execution and main thread can move ahead

			once all the called functions executes wg.done, main thread knows that
			all the go routines are executed and can begin execution of code after
		 	wg.wait()
	*/

	wg.Add(4)
	program_start_time = time.Now()
	fmt.Println("----------------------------------------------------")
	go doSomeWork("A", &wg) //passing waitgroup as a reference (Not a copy)
	go doSomeWork("B", &wg)
	go doSomeWork("C", &wg)
	go doSomeWork("D", &wg)
	wg.Wait()
	fmt.Println("----------------------------------------------------")
	program_end_time = time.Now()
	produceFinalReport()

}

func doSomeWork(work_id string, wg *sync.WaitGroup) {
	const NumOfJobs = 5
	defer wg.Done() //defer means this statement will execute just before the func ends
	for job_id := 1; job_id <= NumOfJobs; job_id++ {
		wait_ms := randomWait()
		fmt.Printf("job%d of work%s took:     %d ms\n", job_id, work_id, wait_ms)
		total_completed_jobs++
		total_worker_busyness += wait_ms
	}
}

func randomWait() int {
	const MinWaitMs = 50
	const MaxWaitMs = 350
	waitMilliseconds := MinWaitMs + rand.Intn(MaxWaitMs-MinWaitMs)
	time.Sleep(time.Duration(waitMilliseconds) * time.Millisecond)
	return waitMilliseconds
}

func produceFinalReport() {
	fmt.Println("         Program time: ", program_end_time.Sub(program_start_time))
	fmt.Println(" Total completed jobs: ", total_completed_jobs)
	fmt.Println("Total worker busyness: ", total_worker_busyness, "ms")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
