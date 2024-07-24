package main

import (
	"fmt"
	"strings"
)

func main() {
	
	//SPLIT
	data := "1,2,3,4,5"
	parts := strings.Split(data, ",")
	fmt.Println("SPLITTED STRING", parts)

	//COUNT
	data_count := strings.Count(data, ",")
	fmt.Println("DATA COUNT :", data_count)

	//TRIM SPACE
	space := "           HELLO SD           "
	strtrim := strings.TrimSpace(space)
	fmt.Println("TRIMMED SPACE: ", strtrim)
	
	//JOIN
	fname := "SHREYASH"
	lname := "DESAI"
	joined := strings.Join([] string {fname,lname}," ")
	fmt.Println("JOINED STRINGS : ",joined)

}
