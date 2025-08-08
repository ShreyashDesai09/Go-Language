package main

import (
	"fmt"
)

func main() {
	// 	fmt.Println("SD")
	// 	a := 10
	// 	b := 20
	// 	fmt.Println(a + b)

	// 	for i := 0 ; i <=10 ; i++ {
	// 		fmt.Println(i)
	// 	}

	// 	var add = []int{}
	// 	add = append(add,10)
	// 	fmt.Println(add[0])

	// a := 10
	// if (a == 5) {
	// 	fmt.Println("ten")
	// } else {
	// 	fmt.Println("FIVE")
	// }

	// scanner := bufio.NewReader(os.Stdin)
	// text, _ := scanner.ReadString('\n')
	// fmt.Println(text)

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Println(num)

	fmt.Println("ENTER PAGES:-")

	Pages := []int{}

	var p int
	fmt.Scan(p)
	// scanner := bufio.NewScanner(os.Stdin)
	
	var z int
	for i := 0; i >= p; i++ {
		fmt.Print("Enter Pages:")
		fmt.Scan(&z)
		Pages = append(Pages, z)
	}

	// var x int
	fmt.Println("Enter Book No. =")



}
