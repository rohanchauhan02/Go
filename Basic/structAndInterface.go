package main

import "fmt"

type cat struct{
	name string
	age int
}
func (c *cat) say(){
	fmt.Println("meow")
}
func main(){
	cat:=cat{
		name:"Mosko",
		age:13,
	}
	
	// cat.say()
	fmt.Println(cat.name,cat.age)
}