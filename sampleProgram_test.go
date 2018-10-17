package main

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestCalculate(t *testing.T) {
	if Calculate(10) != 12 {
		t.Error("Test failed: 2 should be added to 10 and return 12")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input int
		expectedResult int
	} {
		{2, 4},
		{10, 12},
		{90, 92},
		{-12, -10},
		{1000, 1002},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expectedResult {
			t.Errorf("Test Failed: %d was inputted, %d was the expected output, actual output: %d",
				test.input, test.expectedResult, output)
		}
	}
}

func TestFib10(t *testing.T) {
	if Fib(10) != 55 {
		t.Error("Test failed: input of 10 should return 55")
	}
}

func TestFib20(t *testing.T) {
	if Fib(20) != 6765 {
		t.Error("Test failed: input of 20 should return 6765")
	}
}

func TestFib40(t *testing.T) {
	if Fib(40) != 102334155 {
		t.Error("Test failed: input of 40 should return 102334155")
	}
}

func TestFib50(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Fib50 in short mode")
	}
	if Fib(50) != 12586269025 {
		t.Error("Test failed: input of 50 should return 12586269025")
	}
}

// --------------------------------------------------------------------------------

func TestParallel(t *testing.T) {
	t.Parallel()
	TestFib10(t)
	TestFib20(t)
	TestFib40(t)
}

// --------------------------------------------------------------------------------

func BenchmarkFib30(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(30)
	}
}

func BenchmarkFib50(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping Fib50 bench in short mode")
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(50)
	}
}

// --------------------------------------------------------------------------------

func ExampleFib10() {
	fmt.Println(Fib(10))
	// Output: 55
}

func ExamplePerm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
	// Unordered output:
	// 2
	// 1
	// 3
	// 0
	// 4
}



