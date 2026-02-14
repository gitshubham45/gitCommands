package main

import "fmt"

func main() {
	fmt.Println("Topic branch update")
	a := 0
	b := 1
	c := a + b

	d := c + b
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println("Main branch update")
}
