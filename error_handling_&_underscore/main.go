package main

import (
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ERROR")
	}
	return a / b, nil

}

func main() {

	fmt.Println("O/P : ")
	ans, _ := div(10, 0)
	fmt.Println("ANS IS =", ans)

}
