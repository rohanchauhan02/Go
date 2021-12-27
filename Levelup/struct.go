package main

import "fmt"

func main(){
	fmt.Println("Hello")
	type ninja struct{ //just like a class
		name string
		weapons []string
		level int
	}
	wallace:=ninja{name:"wallace"}
	fmt.Println(wallace.name)
	wallace=ninja{
		name:"Rollace", //change name
		weapons:[]string{"ninja star"},
		level:1,
	}
	fmt.Println(wallace)
	wallace.level++
	fmt.Println(wallace.level)

	wallace.weapons=append(wallace.weapons,"Ninja sword")
	fmt.Println(wallace)


	type killer struct{
		title string
		ninja ninja
	}

	rohan:=killer{
		title:"Golang Killer",
		ninja:wallace, // Pass by value
	}
	fmt.Println(rohan)
	fmt.Println(rohan.ninja.level)
	rohan.ninja.level++
	fmt.Println(rohan.ninja.level)
	fmt.Println(wallace.level)


	type addressedDojo struct{
		name string
		ninja *ninja
	}
	addressed:=addressedDojo{
		name:"Addressed Golang Dojo",
		ninja: &wallace,
	}
	fmt.Println(addressed)
	fmt.Println(*addressed.ninja)
	addressed.ninja.level++
	fmt.Println(*&addressed.ninja.level)

	fmt.Println(wallace)
	fmt.Println(wallace.level)

	jonnyPointer:=new(ninja)
	jonnyPointer.name="Jonny"
	jonnyPointer.weapons=[]string{"ninja star","ninja killer","ninja champ"}
	jonnyPointer.level++
	fmt.Println(*jonnyPointer)
	

	intern:=ninjaIntern{name:"Intern",level:1}
	intern.sayHello()
	fmt.Println(intern)
	intern.sayName()


}

type ninjaIntern struct{
	name string
	level int
}
func (ninjaIntern) sayHello(){
	fmt.Println("Hello")
}
func (n ninjaIntern) sayName(){
	fmt.Println(n.name)
}