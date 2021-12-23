package main

import "fmt" 
const (
	x=iota
	_          //skip
	y=iota
	z=iota
)
func main() {
	// a := 4 + 5i
	// c := complex(4, 5)
	// fmt.Printf("%v,%v", real(c), imag(c))
	// primitive data type
	// var b int8 =10     //-128 to 127
	// Number 1. Integer 2. float
	// const i int=20
	// i=99
	var arr=[...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
}
