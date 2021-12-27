package main

import "fmt"

func main(){
	var a int=100
	var b int=200
	fmt.Printf("before swap value of a,b id %d,%d\n",a,b)
	swap(a,b)
	fmt.Printf("after swap value of a,b %d,%d\n",a,b)

}
func swap(x int,y int){
	var temp int
	temp=x
	x=y
	y=temp
}

//no change happens