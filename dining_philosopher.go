package main

import (
	"fmt"
	"sync"
	"time"
)

type fork struct{ sync.Mutex }
type philosopher struct {
	id                  int
	leftfork, rightfork *fork
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.leftfork.Lock()
		p.rightfork.Lock()
		time.Sleep(time.Second)
		say("eating", p.id)
		p.leftfork.Unlock()
		p.rightfork.Unlock()
		say("finish eating", p.id)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func say(msg string, id int) {
	fmt.Printf("philosopher %d is %s\n", id+1, msg)
}

var wg sync.WaitGroup

func main() {
	var count int
	fmt.Println("Enter the number of philosophers:")
	fmt.Scan(&count)
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{id: i, leftfork: forks[i], rightfork: forks[(i+1)%count]}
		wg.Add(1)
		go philosophers[i].eat()
	}
	wg.Wait()
}
