package main

import "fmt"

func main(){
	var a int=100
	var b int=200
	fmt.Printf("before swap value of a id %d\n",a)
	fmt.Printf("before swap value of b id %d\n",b)
	swap(&a,&b)
	fmt.Printf("after swap value of a id %d\n",a)
	fmt.Printf("after swap value of b id %d\n",b)
}
func swap(x *int,y *int){
	var temp int
	temp=*x
	*x=*y
	*y=temp
}