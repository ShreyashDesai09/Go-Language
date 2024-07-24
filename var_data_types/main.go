package main

import "fmt"

//Global Declaration showuld be in
//var var_name data_type
var name string = "shreyash"

func main() {

	//Inside Function Declaration type

	var srname string = "desai"

	//Shortform Declaration Type
	fullName := "SHREYASH DESAI"

	fmt.Println(name)
	fmt.Println(srname)
	fmt.Println(fullName)

	//int , bool

	num := 9
	pool := true

	fmt.Println(fullName)
	fmt.Println(num)
	fmt.Println(pool)
	fmt.Printf("%T : Data type\n", name)
	fmt.Printf("%T : Data type\n", num)
	fmt.Printf("%T : Data type\n", pool)

	/*
		%T = Data Type
		%s = string
		%d = int
		%f = float64
		%v = value
	*/

	/*    NOTE   */
	/*  If we use capital letter to define something it says it is PUBLIC
	    If it is in small letter it is in PRIVATE

		i.e var Shreyash string = can be used for this file as well as other
			var shreyash string = can only be used for this file only
	*/

}
