package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fn string) ([3]float64, error) {
	var numbers [3]float64
	f, err := os.Open(fn)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return numbers, nil
}
func main() {
	prices, err := GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, price := range prices {
		fmt.Println(price)
	}
}

// package main

// import (
// 	"fmt"
// )

// func printPointer(myBoolPointer *bool) {
// 	fmt.Println(*myBoolPointer)
// }
// func main() {
// 	myBool := true
// 	//printPointer(myBool) //cannot use myBool (variable of type *bool in argument to printPointer)
// 	printPointer(&myBool)

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	notes := [3]string{"do", "re", "mi"}
//for i := 0; i <= 3; i++ {
// 	for i := 0; i <= len(notes); i++ {
// 	fmt.Println(i, notes[i])
// }

// for i, note := range notes {
// 	fmt.Println(i, note)
// }

// for note := range notes { //인덱스만 출력
// 	fmt.Println(note)
// }

// 	for _, note := range notes { //_의 역할 : 에러 뜨는거 무시하고 싶을 때
// 		fmt.Println(note)
// 	}
// }
