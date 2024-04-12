package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Printf("v1 %d of type %T\n", v1, v1)

	// non-existent key
	v3 := m["k3"]
	fmt.Println("v3", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// removes all
	clear(m)
	fmt.Println("map:", m)

	// second return value - if value was present or default
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	_, prs2 := m["k20"]
	fmt.Println("prs:", prs2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
