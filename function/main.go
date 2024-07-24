package main

import (
	"fmt"
)

func add(a int, b int) (result int, s string) {

	result = a + b
	s = "ADDITION"
	return
}

func prac(z string, i string) string {
	fmt.Println(z, i)
	return z + i
}

func main() {

	fmt.Println(add(2, 2))
	prac("Shreyash", "Desai")

}
