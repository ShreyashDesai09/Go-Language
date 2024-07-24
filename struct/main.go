package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type detail struct {
	appartment string
	house      int
	flat       int
}

type contact struct {
	mob  int
	mail string
}

type all struct {
	all_person  person
	all_detail  detail
	all_contact contact
}

func main() {

	//Method 1: Struct Declared one by one
	var shreyash person

	shreyash.fname = "shreyash"
	shreyash.lname = "desai"
	shreyash.age = 22
	
	fmt.Println("one by one" ,shreyash)

	//METHOD 2 all together i/p
	Person1 := person{
		fname: "SHREYASH",
		lname: "DESAI",
		age:   2200,
	}

	fmt.Println("---------------------------------------\nALL TOGETHER I/P",Person1)

	//METHOD 3 NEW () I/P
	Person2 := new(person)

	Person2.fname = "S"
	Person2.lname = "D"
	Person2.age = 222000

	fmt.Println("------------------------------\nI/P USING NEW()",Person2)


	var allD all

	allD.all_person.fname = "Shreyash"
	allD.all_person.lname = "desai"
	allD.all_person.age = 22

	allD.all_detail = detail{
		appartment: "Vastushree Gardens",
		house: 3,
		flat: 301,
	}

	allD.all_contact = contact{
		mob: 942629500,
		mail: "shreyashdesai31796@gmail.com",
	}

	fmt.Println("-------------------------\nACCESSING USED STRUCT AND CHANGINFNG DETAILS AND GIVING O/P\n",allD)
}
