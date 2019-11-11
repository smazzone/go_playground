package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1, v1present := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("v1 present:", v1present)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//we ignore the return value for k2 as it is not present
	_, v2present := m["k2"]
	fmt.Println("v2 present:", v2present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
