package main

import "fmt"

func main() {
	fmt.Println(add())
	var ans int =diff(20,99)
	fmt.Println(ans)
	summation(1,2,3,4,5,6,7,8,9)
	var m,d int=multiplication(2,5)
	fmt.Println(m,d)
	fmt.Println(isEven(4))
	num:=4
	//call by value
	callByValue(num)
	fmt.Println(num)
	//call by referance
	num1:=4
	callByReferance(&num1)
	fmt.Println(num1)

	
}

func add()int {
	var a int=20
	var b int=30
	return a + b
}

func diff(x int,y int) int{
	return y-x
}

func summation(vals ...int){
	sum:=0
	for _,n:=range vals{  //Another way to use for loop
		sum+=n
	}
	fmt.Println(sum)
}

func multiplication(x int,y int)(int,int){ //return multiple values 
	return (x*y),(x/y)
}

func isEven(x int) bool{
	return x%2==0
}

//call by value call by referance

func callByValue(num int){
	num=40
	fmt.Println(num)
}

func callByReferance(num1 *int){
	*num1=40
	fmt.Println(*num1)
}
