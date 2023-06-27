package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Itdotaer/learn-golang/distributed_system/pack"
)

func main() {
	var group errgroup.Group
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(group)

	//notSafeCounter()
	//safeCounter()
	//safeCounterWithLock()
	//safeCounterWithAtomic()
	//producerConsumerDemo()

	//factory := func_class.WeaponFactory{}
	//factory.Hit("gun")
}

func producerConsumerDemo() {
	producerConsumer := pack.NewProducerConsumer(10)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var counter int
		for true {
			counter++
			producerConsumer.Produce(fmt.Sprintf("produce num: %d", counter))
		}

		wg.Done()
	}()

	go func() {
		for true {
			producerConsumer.Consume()
		}

		wg.Done()
	}()

	wg.Wait()
}

func notSafeCounter() {
	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}

func safeCounter() {
	var counter int

	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}

func safeCounterWithLock() {
	var counter int
	var wg sync.WaitGroup
	lock := pack.NewLock()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock(time.Second * 2) {
				println("lock failed")
				return
			}
			counter++
			lock.UnLock()
		}()
	}

	wg.Wait()
	println(counter)
}

func safeCounterWithAtomic() {
	var counter uint64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&counter, 1)
		}()
	}
	wg.Wait()
	println(counter)
}
