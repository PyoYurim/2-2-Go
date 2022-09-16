package main

import (
	"fmt"
	"reflect"
)

func main() {
	var quantity int
	//var l, w float64
	//l, w = 1.2, 2.4

	//var l, w float64 = 1.2, 2.4

	//var l, w = 1.2, 2.4

	//var l, w float32 = 1.2, 2.4
	l, w := 1.2, 2.4

	var test bool
	fmt.Println(test)

	fmt.Printf("%.1f %.1f\n", l, w)
	fmt.Println(l*w, "제곱미터")
	quantity = 10
	fmt.Println(quantity)
	fmt.Println(reflect.TypeOf(l))
}
