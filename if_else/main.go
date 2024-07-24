package main

import (
	"fmt"
)

func main() {

	day := "MONDAY"

	if day == "MONDAY" {
		fmt.Println("IT's Week Day")
	} else {
		fmt.Println("It's not weekday")
	}

	var num int
	fmt.Println("Enter 1 2 3 ")
	// scanner := bufio.NewReader(os.Stdin)
	// num , _ := scanner.ReadString('\n')
	fmt.Scan(&num)
	fmt.Println("I/P is ", num)

	if num == 1 {
		fmt.Println("1")
	} else if num == 2 {
		fmt.Println("2")
	} else if num == 3 {
		fmt.Println("3")
	} else {
		fmt.Println("NO NO. BETWEEN 1 2 3")
	}

}
