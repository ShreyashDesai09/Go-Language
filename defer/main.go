package main

import "fmt"

func main() {

	//defer is used to print something at end of the o/p
	//it uses LIFO (Last In First Out) Method like STACK

	fmt.Println("ORDER")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	//I/P taken as 3 2 1
	//O/P given as 1 2 3
}
