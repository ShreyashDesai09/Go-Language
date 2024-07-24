package main

import "fmt"

func main() {

	var num int
	var ptr *int

	num = 20
	ptr = &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)

}