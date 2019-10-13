package main

import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	elements []int
	swapCounter [][2]int
}

func (h *IntHeap) Len() int           { return len(h.elements) }
func (h *IntHeap) Less(i, j int) bool { return h.elements[i] < h.elements[j] }

func (h *IntHeap) Swap(i, j int) {
	h.swapCounter = append(h.swapCounter, [2]int{i, j})
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.elements = append(h.elements, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h.elements
	n := len(old)
	x := old[n-1]
	h.elements = old[0 : n-1]
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	inputSlice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var new int
		fmt.Scan(&new)
		inputSlice = append(inputSlice, new)
	}
	workHeap := &IntHeap{elements: inputSlice}
	heap.Init(workHeap)
	fmt.Println(len(workHeap.swapCounter))
	for _, i := range workHeap.swapCounter{
		fmt.Printf("%d %d\n", i[0], i[1])
	}
}
