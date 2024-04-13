package main

// import "fmt"

func main() {

}

func Repeat(letter string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += letter
	}
	return result
}

func RepeatTimes(letter string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += letter
	}
	return result
}
