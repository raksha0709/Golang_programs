package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup, workduration time.Duration) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(workduration)
	fmt.Printf("worker %d has reached the checkpoint\n", id)
	checkpoint <- true
	<-resume
	fmt.Printf("worker %d is resumimg ...\n", id)
}
func main() {
	var numworkers int
	fmt.Println("enter the number of workers:")
	fmt.Scan(&numworkers)
	var workduration int
	fmt.Println("enter the workduration of workers:")
	fmt.Scan(&workduration)
	checkpoint := make(chan bool)
	resume := make(chan bool)
	var wg sync.WaitGroup
	for i := 1; i <= numworkers; i++ {
		wg.Add(1)
		go workers(i, checkpoint, resume, &wg, time.Duration(workduration)*time.Second)
	}
	for i := 1; i <= numworkers; i++ {
		<-checkpoint
	}
	fmt.Println("all workers have reached the checkpoint")
	fmt.Println("resuming all work")
	for i := 1; i <= numworkers; i++ {
		resume <- true
	}
	wg.Wait()
	fmt.Println("all workers have finished working...")

}
