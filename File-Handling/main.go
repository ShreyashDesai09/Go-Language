package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//METOD I :- TO CREATE FILE

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	defer file.Close()

	//METHOD II :- WRITING IN FILE

	contens := "HELLO THIS IS SHREYASH DESAI"
	_, error := io.WriteString(file, contens+"\n")
	if error != nil {
		fmt.Println("ERROR", error)
		return
	}
	defer file.Close()

	//METHOD III :- 1)READ FROM FILE BUFFER METHOD

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			fmt.Println("ERROR", err)
			return
		}
		fmt.Println("BUFFER TYPE O/P:", buffer[:n])
	}

	//METHOD III :- 1)READ FROM FILE osuntil METHOD

	content, error := ioutil.Readfile("hello.txt")
	if error != nil {
		fmt.Println("ERROR", error)
		return
	}
	fmt.Println(content)

}
