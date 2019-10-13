package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Table struct {
	id int
	parent *Table
	elementsNumber int
}

func NewTable(id, elementsNumber int) Table {
	return Table{
		id:             id,
		elementsNumber: elementsNumber,
	}
}

func (s *Table) Table() *Table {
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
		return s.parent.Table()
	}
}

func (s *Table) ID() int {return s.Table().id}
func (s *Table) ElementsNumber() int {return s.Table().elementsNumber}

func (s *Table) Merge(table *Table) {
	first, second := s.Table(), table.Table()
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
	var tableNumber, queryNumber int
	fmt.Scan(&tableNumber, &queryNumber)
	var max int
	tableSlice := make([]*Table, 0, tableNumber)
	for t := 0; t < tableNumber; t++ {
		var en int
		fmt.Scan(&en)
		if en > max {
			max = en
		}
		nt := NewTable(t+1, en)
		tableSlice = append(tableSlice, &nt)
	}

	reader := bufio.NewReader(os.Stdin)

	for q := 0; q < queryNumber; q++ {
		b, _ := reader.ReadBytes(32)
		first, _ := strconv.ParseInt(string(b[:len(b)-1]), 10, 64)
		b, _, _ = reader.ReadLine()
		second, _ := strconv.ParseInt(string(b), 10, 64)
		tableSlice[first-1].Merge(tableSlice[second-1])
		if sen := tableSlice[first-1].ElementsNumber(); sen > max {
			max = sen
		}
		fmt.Println(max)
	}
}
