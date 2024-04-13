package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4
	assertCorrectResult(t, result, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectResult(t testing.TB, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("expected: '%d', but got '%d'", expected, result)
	}
}
