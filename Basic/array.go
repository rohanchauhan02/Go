package main

import "fmt"

func main() {
	// var arr [4]int =[4]int {1,2,3,4}
	// arr :=[...]int {1,2,3,4,5,6,7,8,9}
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr[1] = 20
	fmt.Println("----------------arr--------------")
	fmt.Println(arr)
	arr2 := arr // element copy
	arr2[3] = 30
	fmt.Println("----------------arr--------------")
	fmt.Println(arr)
	fmt.Println("----------------arr2--------------")
	fmt.Println(arr2)
	// fmt.Println(arr[1:5])
	// fmt.Println(arr[:5])

	//2d array
	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {5, 6, 7}, {0, 9, 8}}
	fmt.Println("----------------matrix--------------")
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])

	////////////////////////////////////////     map   ///////////////////////////////////
	empSal := make(map[string]int) //declaration

	empSal = map[string]int{ //initialization
		"Neha": 20000,
		"Raj":  3000,
		"Atul": 15000,
	}

	fmt.Println("----------------empSal--------------")
	empSal["Rohan"] = 3000000 // Add push key value pair
	// empSal["Rohan"]=35
	fmt.Println((empSal))
	empSal["Atul"] = 5          //Change value
	fmt.Println(empSal["Atul"]) //Get key value pair

	//inline declaration and initialization

	dict := map[string]int{
		"Neha": 20000,
		"Raj":  3000,
		"Atul": 15000,
	}
	fmt.Println("----------------dict-----------------")
	fmt.Println((dict))
	//memory management MAP

	empSal["Rohan"] = 35
	emp := empSal //Copy of Address   //V.V.I
	// emp["Rohan"]=30
	// empSal["Rohan"]=40
	delete(empSal, "Raj")
	fmt.Println("----------------emp---------------------")
	fmt.Println(emp)
	fmt.Println("----------------empSal--------------")
	fmt.Println((empSal))

	fmt.Println("----------------ByDefault value of a non existed Key--------------")
	fmt.Println((empSal["Anuj"]))
	anuj, ok := empSal["Anuj"] //ok is false if "Anuj" is not available
	fmt.Println(anuj, ok)

	//we can use _ to skip variable name declaration
	_, flag := empSal["Puja"] //flag is false if "Puja" is not available and show no error
	fmt.Println(flag)

}
