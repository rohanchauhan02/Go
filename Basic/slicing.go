package main

import "fmt"

func main(){
	fmt.Println("Slicing")

	//array 

	// var arr [4]int=[4]int{1,2,3,4}   
	var arr1 =[...]int{1,2,3,4}      
	fmt.Println(len(arr1))
	fmt.Println("-------------------Slicing----------------")
	//slicing  {Similar to array} 
	var arr[]int=[]int{1,2,3,4}
	
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Println("-------------------Slicing Memory Management----------------")

	arr2:=arr  //In slicing array address is copying to another array
	fmt.Println("------Before-----")
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("------After------")
	arr2[1]=0 
	fmt.Println(arr)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))
	fmt.Println(arr2)
	fmt.Println("-------------------len vs cap----------------")
	// fmt.Println(cap(arr))
	// fmt.Println(cap(arr2))
	var arr3[]int=make([]int,2,4) //another method to declare slice
	fmt.Println("------Before-----")
	fmt.Println(cap(arr3))
	fmt.Println(len(arr3))
	fmt.Println("------After------") // len and cap may or may not be equal
	arr3=append(arr,3000)
	fmt.Println(arr3)
	fmt.Println(arr) // No change in arr
	fmt.Println(cap(arr3))
	fmt.Println(len(arr3))
	fmt.Println("-------------------append and len vs cap----------------")
	arr4:=append(arr,arr3...)
	fmt.Println(arr4)
	fmt.Println(cap(arr4)) //Why capacity is 10?????
	fmt.Println(len(arr4))

}