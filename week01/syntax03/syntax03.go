package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("숫자 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, _ := rd.ReadString('\n')
	fmt.Println(in)

}
