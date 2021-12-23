package main
import "fmt"

type Employee struct{ //global
	EmpId int
	EmpName string
	EmpMobile []string
	// DeptName string
}
func main(){
	fmt.Println("Stuct")
	emp:=Employee{
		//not best way to assign
		// 101,
		// "Rohan",
		// []string{"9113769631","9205916092"},

		//best way to assign 
		EmpId:101,
		EmpName:"Rohan",
		EmpMobile:[]string{"9113769631","9205916092"},
	}
	fmt.Println(emp)

	//Memory Management
	fmt.Println("----------Memory Management-----------")
	fmt.Println("---Method1---")
	emp2:=emp          //Elements copy to emp2
	emp.EmpName="Mohan"
	emp2.EmpName="Sohan"
	fmt.Println(emp)
	fmt.Println(emp2)
	fmt.Println("---Method2---")
	emp3:=&emp          //Copy address of emp
	emp.EmpName="Mohan"
	// emp3.EmpName="Sohan"
	fmt.Println(emp)
	fmt.Println(*emp3)

	
}