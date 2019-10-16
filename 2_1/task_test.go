package main

import "testing"

var correctResultExamples = map[string]int{
	"[]": 0,
	"{}[]": 0,
	"[()]": 0,
	"(())": 0,
	"{[]}()": 0,
	"{": 1,
	"{[}": 3,
	"foo(bar);": 0,
	"foo(bar[i);": 10,
}

func TestAll(t *testing.T){
	for example, expected := range correctResultExamples{
		if result := checkBracket(example); result != expected {
			t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, expected)
		}
	}
}
