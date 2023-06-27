package pack

import "time"

type ProducerConsumer struct {
	c chan string
}

func NewProducerConsumer(num int) ProducerConsumer {
	var producerConsumer ProducerConsumer
	producerConsumer.c = make(chan string, num)
	return producerConsumer
}

func (pc ProducerConsumer) Produce(produce string) {
	println("produce-->", produce)
	pc.c <- produce
	time.Sleep(time.Second * 1)
}

func (pc ProducerConsumer) Consume() {
	select {
	case produce := <-pc.c:
		println("consume-->", produce)
	}
	time.Sleep(time.Second * 2)
}
