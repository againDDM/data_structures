package main

import (
	"container/heap"
	"fmt"
	"os"
)

type Processor struct {
	busy int
	index int
}

func (proc Processor) String() string {
	return fmt.Sprintf("%d %d", proc.index, proc.busy)
}

type CPUqueue []*Processor

func newCPUqueue(size int) CPUqueue {
	result := make([]*Processor, size)
	for index, _ := range(result){
		result[index] = &Processor{
			busy:  0,
			index: index,
		}
	}
	return result
}

func (pq CPUqueue) Len() int      { return len(pq) }
func (pq CPUqueue) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}

func (pq CPUqueue) Less(i, j int) bool {
	if pq[i].busy == pq[j].busy {
		return pq[i].index < pq[j].index
	}
	return pq[i].busy < pq[j].busy
}

func (pq *CPUqueue) Push(x interface{}) {
	processor := x.(*Processor)
	*pq = append(*pq, processor)
}

func (pq *CPUqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *CPUqueue) rotate(busy int) Processor {
	proc := (*pq)[0]
	result := (*proc)
	proc.busy += busy
	heap.Fix(pq, 0)
	return result
}


func main() {
	var procNumber, taskNumber int
	fmt.Fscan(os.Stdin, &procNumber, &taskNumber)
	CP := newCPUqueue(procNumber)
	for task := 0; task < taskNumber; task++ {
		var new int
		fmt.Scan(&new)
		fmt.Println(CP.rotate(new))
	}
}