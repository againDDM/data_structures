package main

import (
	"fmt"
	"os"
)

type limitedQueue struct {
	limit uint64
	queue []int
}

func newLimitedQueue(limit uint64) limitedQueue {
	return limitedQueue{
		limit: limit,
		queue: make([]int, 0, limit),
	}
}

func (q *limitedQueue) len() uint64 {
	return uint64(len(q.queue))
}

func (q *limitedQueue) isFull() bool {
	return q.len() == q.limit
}

func (q *limitedQueue) isEmpty() bool {
	return q.len() == 0
}

func (q *limitedQueue) promote(till int) {
	var numPop int
	for i, packet := range q.queue {
		if packet > till {
			break
		} else {
			numPop = i + 1
		}
	}
	q.queue = q.queue[numPop:]
	//	fmt.Println(q.queue)
}

func (q *limitedQueue) last() int {
	if q.isEmpty() {
		return 0
	}
	return q.queue[q.len()-1]
}

func (q *limitedQueue) enqueue(arrival, duration int) (result int) {
	q.promote(arrival)
	if q.isFull() {
		return -1
	}
	if q.last() < arrival {
		result = arrival
	} else {
		result = q.last()
	}
	q.queue = append(q.queue, result+duration)
	return
}

func main() {
	var size, n int
	fmt.Fscan(os.Stdin, &size, &n)
	workQueue := newLimitedQueue(uint64(size))
	for i := 0; i < n; i++ {
		var arrival, duration int
		fmt.Fscan(os.Stdin, &arrival, &duration)
		fmt.Println(workQueue.enqueue(arrival, duration))
	}
}
