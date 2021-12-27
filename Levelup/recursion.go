package main

import (
	"fmt"
	"math"
)

func printIncreasing(a int, b int) {
	if b == a {
		return
	}
	printIncreasing(a, b-1)
	fmt.Println(b)
}

func printDecreasing(a int, b int) {
	if a == b {
		return
	}
	printDecreasing(a+1, b)
	fmt.Println(a)

}
func oddEven(a int, b int) {
	if a == b {
		return
	}
	if a%2 != 0 {
		fmt.Println(a)
	}
	oddEven(a+1, b)
	if a%2 == 0 {
		fmt.Println(a)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func displayArray(arr []int, idx int) {
	if idx == -1 {
		return
	}
	displayArray(arr, idx-1)
	fmt.Println(arr[idx])
}
func maximum(arr []int, idx int) float64 {
	if idx == len(arr) {
		return math.MinInt64
	}
	return math.Max(maximum(arr, idx+1), float64(arr[idx]))

}
func minimum(arr []int, idx int) float64 {
	if idx == len(arr) {
		return math.MaxInt64
	}
	return math.Min(minimum(arr, idx+1), float64(arr[idx]))
}

func find(arr []int, idx int, data int) bool {
	if idx == len(arr) {
		return false
	}
	if arr[idx] == data {
		return true
	}

	return find(arr, idx+1, data)
}
func firstIndex(arr []int, idx int, data int) int {
	if idx == len(arr) {
		return -1
	}
	if arr[idx]==data{
		return idx
	}
	return firstIndex(arr, idx+1, data)
}

func lastIdx(arr []int,idx int,data int) int{
	if idx==len(arr){
		return -1
	}
	res:=lastIdx(arr,idx+1,data)
	if res!=-1{
		return res
	}
	if arr[idx]==data {
		return idx
	}else{
		return -1
	}
	// ternary operator is not included in golang
}
func allIndex(arr []int,idx int,data int,count int) []int{
	if idx==len(arr){
		return make([]int,count,count)
	}
	if arr[idx]==data{
		count++
	}
	ans:=allIndex(arr,idx+1,data,0)
	if arr[idx]==data{
		ans[count]=idx
	}
	return ans
}

func main() {
	// fmt.Println("Hello")
	// printIncreasing(0,10)
	// printDecreasing(0,10)
	// oddEven(0,10)
	// fmt.Println(factorial(5))
	var arr []int
	arr = []int{51, 2, 30, 4, 5, 60}
	// displayArray(arr, 5)
	var val1 float64 = maximum(arr, 0)
	fmt.Println(val1)
	var val2 float64 = minimum(arr, 0)
	fmt.Println(val2)
	fmt.Println(find(arr, 0, 40))
	arr1 := []int{51, 2, 30, 4, 5, 60, 4, 5, 3, 7, 9, 4, 50, 4, 5}

	fmt.Println(firstIndex(arr1, 0, 4))
	fmt.Println(lastIdx(arr1,0,4))
	fmt.Println(allIndex(arr1,0,4,0))
}
