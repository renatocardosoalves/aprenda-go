package main

import "fmt"

type hotdog int

var a hotdog = 10

func main() {
	b := 20
	a = hotdog(b)

	fmt.Println(a)
}
