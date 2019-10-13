package main

import 	(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Set struct {
	id int
	parent *Set
	elementsNumber int `default:"1"`
}

func NewSet(id, elementsNumber int) Set {
	return Set{
		id:             id,
		elementsNumber: elementsNumber,
	}
}

func (s *Set) Set() *Set {
	switch {
	case s.parent == nil:
		return s
	case s.parent.parent == nil:
		return s.parent
	case s.parent.parent.parent == nil:
		s.parent.elementsNumber -= s.elementsNumber
		s.parent = s.parent.parent
		return s.parent
	default:
		return s.parent.Set()
	}
}

func (s *Set) ID() int             {return s.Set().id}
func (s *Set) ElementsNumber() int {return s.Set().elementsNumber}

func (s *Set) Merge(table *Set) {
	first, second := s.Set(), table.Set()
	if first.id != second.id {
		if first.elementsNumber < second.elementsNumber {
			second.elementsNumber += first.elementsNumber
			first.parent = second
		} else {
			first.elementsNumber += second.elementsNumber
			second.parent = first
		}
	}
}


func main()  {
	var elNumber, eqNumber, neNumber int
	fmt.Scan(&elNumber, &eqNumber, &neNumber)

	sets := make([]*Set, 0, elNumber)
	for el := 0; el < elNumber; el++ {
		ns :=  NewSet(el+1, 1)
		sets = append(sets, &ns)
	}

	reader := bufio.NewReader(os.Stdin)

	for eq := 0; eq < eqNumber; eq++ {
		b, _ := reader.ReadBytes(32)
		first, _ := strconv.ParseInt(string(b[:len(b)-1]), 10, 64)
		b, _, _ = reader.ReadLine()
		second, _ := strconv.ParseInt(string(b), 10, 64)
		sets[first-1].Merge(sets[second-1])
	}

	var success bool = true
	for ne :=0; ne < neNumber; ne++ {
		b, _ := reader.ReadBytes(32)
		first, _ := strconv.ParseInt(string(b[:len(b)-1]), 10, 64)
		b, _, _ = reader.ReadLine()
		second, _ := strconv.ParseInt(string(b), 10, 64)
		if sets[first-1].ID() == sets[second-1].ID() {
			success = false
		}
	}

	if success {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
