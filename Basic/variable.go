package main

import "fmt"

// 2. Global --(Pascal Case means First letter is in Caps )
var Val2 int = 3000

//3. Package_Level --(Camel Case)
var myVal2 = 30001

func main() {
	// fmt.Println("Hello world!")

	//method 1
	var a int //declaration
	a = 55    //initialization
	fmt.Println(a)

	// method 2
	var b int = 500 //declaration with initailization
	var c = 333     //declaration with initailization

	fmt.Println(b)
	fmt.Println(c)

	//method 3
	d := 1111 //less verbose declaration with initailization
	fmt.Println(d)

	// variable scope
	//1. Local
	//2. Global
	//3.Package_Level

	//Local
	var val1 int = 2000
	fmt.Println(val1)
	fmt.Println(Val2)
	fmt.Println(myVal2)
}
