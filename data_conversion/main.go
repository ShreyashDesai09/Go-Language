package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := "SHREYASH"
	numName := 123
	strFloat := "123.123"

	num, _ := strconv.Atoi(name)
	n := strconv.Itoa(numName)
	flt, _ := strconv.ParseFloat(strFloat, 64)

	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", num)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", flt)

}
