package main

import (
	"fmt"
)

func main() {

	nums := [6] int {1,2,3,4,5}
	fmt.Println("ARRAY - ",nums)
	fmt.Println("TOTAL LENGTH -",len(nums))
	fmt.Println("CAPACITY -",cap(nums))
	fmt.Printf("TYPE - %T" , nums)
	for i , val := range nums {
		fmt.Printf("%d  %d \n" , i , val)
	}

}