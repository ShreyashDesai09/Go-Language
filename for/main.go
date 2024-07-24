package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	for index, value := range nums {
		fmt.Printf("%d - %d \n",index,value)
	}

}