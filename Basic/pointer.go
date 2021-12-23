package main 

import "fmt"


type Employee struct{
	EmpName string
}
func main(){
	// var x int=5
	// var b *int=&x  // b holds address of x and *b dereferancing it
	// fmt.Println(x,*b,b) 

	var emp *Employee
	fmt.Println(emp)
	emp= new(Employee) //create object
	fmt.Println(*emp)
}