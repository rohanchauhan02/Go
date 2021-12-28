package main 

import "fmt"

func variadic(params ...int){
	fmt.Printf("%T\n %v\n",params,params)
}


func main(){
	variadic(1,2,3,4,5,6,7,8,9)
}