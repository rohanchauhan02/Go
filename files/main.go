package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file -WebXSolutions"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myfile.txt")
}
func readFile(filename string){
	databyte,err:=ioutil.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("Text data inside the file is \n",databyte) // dtat in binary stream
	fmt.Println("Text data inside the file is \n",string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}