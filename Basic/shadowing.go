package main

import(
	"fmt"
	"strconv"
)

//Shadowing
var val int = 100

// val:=100 //wrong statement it doesnt work on global scope
func main() {
	// var val int = 55
	// fmt.Print(val)
	// fmt.Printf("%v,%T",val,val)

	//casting]
	// var a int32=100
	// var b int64=int64(a)
	// fmt.Printf("%v,%T",b,b)
	// fmt.Printf("Hello : %v ",val)

	// c:=5.5
	// var d float64=float64(c)
	// fmt.Printf("%v,%T",d,d)

	var str int=65
	var e string=strconv.Itoa(str) //convert it to string
	// var e string=string(str) // Return ASCII value
	fmt.Printf("%v,%T",e,e)
}
