package main

import (
	"fmt"
	"os"
	"sync"
	"unicode/utf8"
)

type stackElement struct {
	symbol   rune
	position int
}

type stack struct {
	elements []stackElement
	mux      sync.Mutex
}

func newStack(capacity uint64) stack {
	return stack{
		elements: make([]stackElement, 0, capacity),
	}
}

func (s *stack) push(new stackElement) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.elements = append(s.elements, new)
}

func (s *stack) len() int {
	return len(s.elements)
}

func (s *stack) isEmpty() bool {
	return s.len() == 0
}

func (s *stack) top() (top stackElement, empty bool) {
	empty = s.isEmpty()
	if !empty {
		top = s.elements[s.len()-1]
	}
	return
}

func (s *stack) pop() (top stackElement, empty bool) {
	top, empty = s.top()
	if !empty {
		s.mux.Lock()
		defer s.mux.Unlock()
		s.elements = s.elements[:s.len()-1]
	}
	return
}

func checkBracket(input string) (output int) {
	symbols := utf8.RuneCountInString(input)
	bracketMap := map[rune]rune{
		125: 123, // '}': '{'
		93:  91,  // ']': '['
		41:  40,  // ')': '('
	}
	workStack := newStack(uint64(symbols))
	for i, r := range input {
		switch r {
		case 123, 91, 40: // { [ (
			workStack.push(stackElement{r, i})
		case 125, 93, 41: // } ] )
			var top stackElement
			var empty bool
			top, empty = workStack.pop()
			if empty || bracketMap[r] != top.symbol {
				output = i + 1
				return
			}
		}
	}
	if output == 0 {
		lees, empty := workStack.top()
		if !empty {
			output = lees.position + 1
		}
	}
	return
}

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)
	output := checkBracket(input)
	if output == 0 {
		fmt.Print("Success")
	} else {
		fmt.Println(output)
	}
}
