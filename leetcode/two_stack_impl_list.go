package leetcode

import "github.com/golang-collections/collections/stack"

type CQueue struct {
	from *stack.Stack
	to   *stack.Stack
}

func Constructor() CQueue {
	return CQueue{
		from: stack.New(),
		to:   stack.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.from.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.to.Peek() == nil {
		for this.from.Peek() != nil {
			this.to.Push(this.from.Pop())
		}
	}

	pop := this.to.Pop()
	if pop == nil {
		return 0
	}

	return pop.(int)
}
