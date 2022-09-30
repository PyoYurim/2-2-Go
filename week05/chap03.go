package main

import "fmt"

func main() {
	a := 4
	pa := &a

	fmt.Println(*pa)
	fmt.Println(pa)
	fmt.Println(&a)
}
