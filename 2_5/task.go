package main

import (
	"fmt"
)

type stack struct {
	elements []int64
	maximum  []int64
}

func newStack(capacity int) stack {
	return stack{
		elements: make([]int64, 0, capacity),
		maximum:  make([]int64, 0, capacity),
	}
}

func (s *stack) push(new int64) {
	max, empty := s.max()
	s.elements = append(s.elements, new)
	switch {
	case empty:
		s.maximum = append(s.maximum, new)
	case new >= max:
		s.maximum = append(s.maximum, new)
	default:
		s.maximum = append(s.maximum, max)
	}
}

func (s *stack) len() int {
	return len(s.elements)
}

func (s *stack) isEmpty() bool {
	return s.len() == 0
}

func (s *stack) max() (max int64, empty bool) {
	empty = s.isEmpty()
	if !empty {
		max = s.maximum[s.len()-1]
	}
	return
}

func (s *stack) top() (top int64, empty bool) {
	empty = s.isEmpty()
	if !empty {
		top = s.elements[s.len()-1]
	}
	return
}

func (s *stack) pop() (empty bool) {
	empty = s.isEmpty()
	if !empty {
		newLen := s.len() - 1
		s.elements = s.elements[:newLen]
		s.maximum = s.maximum[:newLen]
	}
	return
}

type queueDSM struct {
	inputStack stack
	outputStack stack
}

func newQueueDSM(size int) queueDSM {
	return queueDSM{
		inputStack:  newStack(size),
		outputStack: newStack(size),
	}
}

func (q *queueDSM) enqueue(new int64) {
	q.inputStack.push(new)
}

func (q *queueDSM) shifting() {
	for {
		next, stop := q.inputStack.top()
		if stop {
			break
		}
		q.outputStack.push(next)
		q.inputStack.pop()
	}
}

func (q *queueDSM) isEmpty() bool {
	return q.inputStack.isEmpty() && q.outputStack.isEmpty()
}

func (q *queueDSM) dequeue() (empty bool) {
	if q.isEmpty() {
		empty = true
		return
	}
	if q.outputStack.isEmpty(){
		q.shifting()
	}
	q.outputStack.pop()
	return
}

func (q *queueDSM) next() (next int64, empty bool) {
	if q.isEmpty() {
		empty = true
		return
	}
	if q.outputStack.isEmpty(){
		q.shifting()
	}
	next, _ = q.outputStack.top()
	return
}

func (q *queueDSM) max() (max int64, empty bool) {
	inMax, inEmpty := q.inputStack.max()
	outMax, outEmpty := q.outputStack.max()
	switch {
	case inEmpty && outEmpty:
		empty = true
	case inEmpty:
		max = outMax
	case outEmpty:
		max = inMax
	case inMax > outMax:
		max = inMax
	default:
		max = outMax
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	workSlice := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		var new int64
		fmt.Scan(&new)
		workSlice = append(workSlice, new)
	}
	var m int
	fmt.Scan(&m)
	workQueue := newQueueDSM(m)
	for i := 0; i < m-1; i++ {
		workQueue.enqueue(workSlice[i])
	}
	for i := m-1; i < n; i++ {
		workQueue.enqueue(workSlice[i])
		new, _ := workQueue.max()
		fmt.Print(new, " ")
		workQueue.dequeue()
	}
}
