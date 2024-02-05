package main

import "testing"

// why it is showing 66.7% coverage for commented below code.. because every possible input must be tested of a function.

//below only i tested even case only

// 66.7 % coverage code

// func TestEven(t *testing.T) {

// 	if Even(4) == false {
// 		t.Error("Expected even number")
// 	}

// }

// This will returns 100% coverage  as i have tested for Even and ODD possiblity both.

func TestEven(t *testing.T) {

	if Even(4) != true {
		t.Error("Expected even number")
	}

	if Even(5) != false {
		t.Error("Expected even number")
	}

}

//Collectively testing

func TestEvenBulk(t *testing.T) {

	tests := []struct {
		input    int
		expected bool
	}{
		{4, true},
		{5, false},
		{8, true},
		{7, false},
	}

	for _, test := range tests {

		output := Even(test.input)

		if output != test.expected {
			t.Error("Test failed {} input , {} expected , receviced : {}", test.input, test.expected, output)
		}
	}
}
