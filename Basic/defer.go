package main
 
import "fmt"

func main(){

	// defer fmt.Println("I") // wait till last
	// fmt.Println("love")
	// fmt.Println("India")
	
	// defer fmt.Println("I") 
	// defer fmt.Println("love")
	// defer fmt.Println("India")

	//uses
	defer fmt.Println("Db connection Closing") 
	fmt.Println("Connection Opening")
	fmt.Println("Database manupulation")
}