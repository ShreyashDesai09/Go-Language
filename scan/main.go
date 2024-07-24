package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var name string

	// fmt.Printf("ENTER WHAT TO BE SCANNED : ")
	// fmt.Scan(&name)
	// fmt.Printf("\nNAME IS : %s", name)

	// above method can scan only one element
	// I.E if we give space it exits scanning
	// --------------------------------------------------------------------------------------
	//BUFIO METHOS

	fmt.Println("\nENTER FULL NAME")
	scanner := bufio.NewReader(os.Stdin)
	fullName, _ := scanner.ReadString('\n')
	fmt.Println("FULL NAME IS :", fullName)

}
