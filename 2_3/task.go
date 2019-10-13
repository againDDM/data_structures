package main

import (
	"fmt"
	"os"
)

type QueueLimited struct {
	limit uint64
	queue []int
}

func NewQueueLimited(limit uint64) QueueLimited {
	return QueueLimited{
		limit: limit,
		queue: make([]int, 0, limit),
	}
}

func (q *QueueLimited) len() uint64 {
	return uint64(len(q.queue))
}

func (q *QueueLimited) isFull() bool {
	return q.len() == q.limit
}

func (q *QueueLimited) isEmpty() bool {
	return q.len() == 0
}

func (q *QueueLimited) promote(till int) {
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

func (q *QueueLimited) last() int {
	if q.isEmpty() {
		return 0
	}
	return q.queue[q.len()-1]
}

func (q *QueueLimited) Enqueue(arrival, duration int) (result int) {
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
	workQueue := NewQueueLimited(uint64(size))
	for i := 0; i < n; i++ {
		var arrival, duration int
		fmt.Fscan(os.Stdin, &arrival, &duration)
		fmt.Println(workQueue.Enqueue(arrival, duration))
	}
}
