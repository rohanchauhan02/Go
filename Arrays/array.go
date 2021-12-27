package main 

import "fmt"

func main(){
	// fmt.Println("hello")
	// arr:=[...]int{1,2,3,4,5,6}
	var arr=[]int{1,2,3,4,5}
	fmt.Println("Array:================")
	displayArray(arr)
	fmt.Println(arr)


	var a []int
	a=make([]int,4,5)
	fmt.Println("Slice:================")
	displaySlice(a)
	fmt.Println(a)
	// fmt.Println(cap(a))
	// a[0]=1
	// fmt.Println(cap(a))
	// a=append(a,5)
	// fmt.Println(cap(a))
	// a=append(a,5)
	// fmt.Println(cap(a))
	// a=append(a,5)
	// fmt.Println(cap(a))    // slice is juct like arrayList its capacity increases by *2 when we append elements more than capacity
	// a=append(a,5)
	// fmt.Println(cap(a))
	// a=append(a,5)
	// a=append(a,5)
	// a=append(a,5)
	// a=append(a,5)
	// a=append(a,5)
	// fmt.Println(cap(a))
	// fmt.Println(a)
	// fmt.Println(cap(a))
	// b:=a
	// b[0]=15				// Copy of address happen
	// fmt.Println(a)
	// fmt.Println(b)
	
	// a:=arr
	// fmt.Println(a)
	// arr[4]=20
	// fmt.Println(arr)

	// var arr1 [3]int= [3]int {1,2,3}
	// arr1[2]=50
	// fmt.Println(arr1)
	// displayArray(arr)

	const count =4
	var base =new([count]int) //return a pointer
	base[1]=10
	fmt.Println(base)
	fmt.Println(*base)
}

func displaySlice(arr []int){
	
	arr[0]=15
	fmt.Println(arr)
}
func displayArray(arr []int){
	// arr[0]=20
	fmt.Println(arr)
}