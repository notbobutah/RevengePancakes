package main

import (
	"testing";
)
var flipTests = []struct {
	stack    string // input
	expected int // expected result
}{
	{"-", 1},
	{"+", 0},
	{"-+", 1},
	{"+-", 2},
	{"+++", 0},
	{"--+-", 3},
	{"--++--", 3},
	{"--+++--", 3},
	{"+++---+++", 2},
	{"-------", 1},
	{"---+++--++--", 5},
	{"--++--++--++--++--++-", 11},
	{"++--++--++--++--++-+", 10},
	{"++--++--++--++--++-+++--++--++--++--++-+", 20},
	{"----------------------------------------------------------", 1},
	{"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", 0},
}

func TestFlip(t *testing.T) {
	title()
	for _, tt := range flipTests {
		actual :=  flip(tt.stack)
		if actual != tt.expected {
			t.Errorf("flip(%s): expected %d, actual %d", tt.stack, tt.expected, actual)
		} else {
			println("Stack test passed for :"+tt.stack);
		}
	}
}
