package main

import "fmt"

func main(){
	type fahrenheit int
	type calcius int
	var f fahrenheit =32
	var c calcius=0
	c=calcius((f-32)*5/9)
	fmt.Println(f,c)
}