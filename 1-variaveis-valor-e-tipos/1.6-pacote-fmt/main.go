package main

import "fmt"

func main() {
	a := "Hello!\nGood morning.\nHow are you?"
	fmt.Println(a)

	b := `Hello!
Good morning.
How are you?`
	fmt.Println(b)

	c, d, e := "Hello!", "Good morning.", "How are you?"
	f := fmt.Sprintln(c, d, e)
	fmt.Print(f)

	fmt.Printf("%s %s %s\n", c, d, e)

	g := fmt.Sprintf("%s %s %s\n", c, d, e)
	fmt.Println(g)
}
