package main

import "testing"

var correctResultExamples = []struct {
	number string
	tree   string
	length int
}{
 {"5", "4 -1 4 1 1", 3},
 {"5", "-1 0 4 0 3", 4},
}

func TestAll(t *testing.T) {
	for _, example := range correctResultExamples{
		if result, err := work(example.number, example.tree); err != nil {
			t.Errorf("Error: %v", err)
		} else if result != example.length {
			t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, example.length)
		}
	}
}
