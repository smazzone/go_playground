package main

import "fmt"

func main() {
	var name string = "Stefano"
	var age int32 = 52

	fmt.Println(name, age)
	fmt.Printf("%T\n", age)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}
