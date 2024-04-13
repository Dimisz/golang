package main

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatTimes("q", 10)
	}
}
func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"
	assertCorrectResult(t, result, expected)
}

func TestRepeatTimes(t *testing.T) {
	result := RepeatTimes("q", 10)
	expected := "qqqqqqqqqq"
	assertCorrectResult(t, result, expected)
}

func assertCorrectResult(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected: '%s', but got: '%s'", expected, result)
	}
}
