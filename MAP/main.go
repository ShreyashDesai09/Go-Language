package main

import "fmt"

func main() {

	studentGrade := make(map[string]int)
	// studentGrade = append(studentGrade[string],"SD")

	//ADDING TO MAP
	studentGrade["SD"] = 9
	studentGrade["VP"] = 8
	studentGrade["SZ"] = 8

	//TO DELETE A MAP ELEMENT
	delete(studentGrade,"VP")

	//Printing using for range
	for index,value := range studentGrade {
		fmt.Printf("%s - %d\n",index,value)
	}

	fmt.Println(studentGrade)

	//To check if MAP value is present or not
	grade , exist := studentGrade["SD"]
	fmt.Println(grade)
	fmt.Println(exist)
}