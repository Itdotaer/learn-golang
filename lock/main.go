package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	var (
		counter1 = 0
		counter2 = 1
	)

	g := sync.WaitGroup{}
	g.Add(2)

	go func() {
		for true {
			lock.Lock()

			counter1 += 2
			println(fmt.Sprintf("one: %d", counter1))
			time.Sleep(time.Second)

			lock.Unlock()
		}

		g.Done()
	}()

	go func() {
		for true {
			lock.Lock()

			counter2 += 2
			println(fmt.Sprintf("two: %d", counter2))
			time.Sleep(time.Second)

			lock.Unlock()
		}

		g.Done()
	}()

	g.Wait()
}
