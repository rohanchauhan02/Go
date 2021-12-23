package main

import "fmt"

func main() {
	// i:=15
	switch i := 6; i {
	// switch {
	case 5, 6, 7:
		// case i>0 && i<=4:
		fmt.Println("line number 11, ", i)
		// fallthrough
	case 1, 2, 3:
		// case i>4:
		fmt.Println("line number 14, ", i)
		// break
		fallthrough //next case will also run
	case 9:
		// case i>10 && i<20:
		fmt.Println("line number 19, ", i)
	default:
		fmt.Println("line number 21, ", i)
	}

	//type

	var x interface{} = "ProngRank"
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
	}

}
