package main

import "fmt"

type Node struct{
	data int
	left *Node
	right *Node
}
var idx int
func (node *Node) create(arr []int){
	if 
}
func printTree(root *Node){
	if root==nil{
		return
	}
	if root.left!=nil{
		fmt.Print(root.left.data)
	}else{
		fmt.Print(".")
	}
	fmt.Print("=>",root.data,"<=")
	if root.right!=nil{
		fmt.Println(root.right.data)
	}else{
		fmt.Println(".")
	}
	printTree(root.left)
	printTree(root.right)
}
func main(){
	tree:=&Node{}
	fmt.Print(tree)
}