package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5}
	fmt.Println("SLICE :",num)
	fmt.Println("LENGTH :" , len(num))
	fmt.Println("Capacity :",cap(num))
	
	num = append(num, 6,7,8,9,10)
	fmt.Println("SLICE :",num)
	fmt.Println("LENGTH :" , len(num))
	fmt.Println("Capacity :",cap(num))

	//SLICES USING MAKE 
	// MAKE([] INT , LENGTH_OF_SLICE , CAPACITY_OF_SLICE)
	fmt.Println("SLICE USING MAKE:-")
	ints := make([] int ,3 , 5)
	fmt.Println("SLICE:",ints)
	fmt.Println("LENGTH:",len(ints))
	fmt.Println("CAPACITY:",cap(ints))
	ints = append(ints, 1,2,3,4,5,6,7)
	fmt.Println("SLICE:",ints)
	fmt.Println("LENGTH:",len(ints))
	fmt.Println("CAPACITY:",cap(ints))

}