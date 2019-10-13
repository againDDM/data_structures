package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type stack struct {
	elements []int64
	maximum  []int64
	mux      sync.Mutex
}

func newStack(capacity int) stack {
	return stack{
		elements: make([]int64, 0, capacity),
		maximum:  make([]int64, 0, capacity),
	}
}

func (s *stack) push(new int64) {
	s.mux.Lock()
	defer s.mux.Unlock()
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
		s.mux.Lock()
		defer s.mux.Unlock()
		newLen := s.len() - 1
		s.elements = s.elements[:newLen]
		s.maximum = s.maximum[:newLen]
	}
	return
}

func main() {
	var q int
	fmt.Fscan(os.Stdin, &q)
	workStack := newStack(q)
	myReader := bufio.NewReader(os.Stdin)
	for i := 0; i < q; i++ {
		command, _, _ := myReader.ReadLine()
		switch {
		case command[0] == 109: // max
			max, _ := workStack.max()
			fmt.Println(max)
		case command[0] == 112 && command[1] == 111: // pop
			workStack.pop()
		default: // push $argument
			argument, _ := strconv.ParseInt(string(command[5:]), 10, 64)
			workStack.push(argument)
		}
	}
}
