package main

import "fmt"

func main() {
	a := 10
	b := "Good morning"

	fmt.Printf("value a = %d, type = %T\n", a, a)
	fmt.Printf("value b = %s, type = %T\n", b, b)

	c, d := 20, "Good night"

	fmt.Printf("value c = %d, type = %T\n", c, c)
	fmt.Printf("value d = %s, type = %T\n", d, d)
}
