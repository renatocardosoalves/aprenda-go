package main

import "fmt"

func main() {
	a, b, c, d := 1, 1.2, "Hello", true

	fmt.Printf("value a = %d, type = %T\n", a, a)
	fmt.Printf("value b = %f, type = %T\n", b, b)
	fmt.Printf("value c = %s, type = %T\n", c, c)
	fmt.Printf("value d = %t, type = %T\n", d, d)
}
