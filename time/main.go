package main

import (
	"fmt"
	"time"
)

func main() {
	//INTIALIZING TIME FUNCTION
	time := time.Now()
	fmt.Println("TIME NOW :", time)

	//FORMATTED TIME
	formatted := time.Format("2006/01/02 Monday 15:04:03 3:04 PM")
	fmt.Println("FORMATTED TIME :", formatted)

	//TIME PARSE
	layout_date := "01/02/2006"
	layout_time := "07/24/2024"
	final, _ := time.Parse(layout_date, layout_time)
	fmt.Println("PARSED TIME", final)
}
