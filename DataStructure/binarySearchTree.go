package main
import "fmt"

//Node 
type Node struct{
	key int
	left *Node
	right *Node
}



//Insert

func (n *Node) Insert(k int){
	if n.key<k{
		//move right
		if n.right==nil{
			n.right=&Node{key:k}
		}else {
			n.right.Insert(k)
		}
	} else if n.key>k{
		//move left
		if n.left==nil{
			n.left=&Node{key:k}
		}else {
			n.left.Insert(k)
		}
	}
}


//Search
func (n* Node) find (k int) bool{
	if n==nil{
		return false
	}
	if n.key==k{
		return true
	}else if n.key>k{
		//move left
		return n.left.find(k)
	}else {
		//move right
		return n.right.find(k)
	}
	
}

func printTree(root *Node){
	if root==nil{
		return
	}
	if root.left!=nil{
		fmt.Print(root.left.key)
	}else{
		fmt.Print(".")
	}
	fmt.Print("=>",root.key,"<=")
	if root.right!=nil{
		fmt.Println(root.right.key)
	}else{
		fmt.Println(".")
	}
	printTree(root.left)
	printTree(root.right)
}


func main(){
	tree:=&Node{key:100}
	tree.Insert(23)
	tree.Insert(33)
	tree.Insert(73)
	tree.Insert(1022)
	tree.Insert(102)
	printTree(tree)
	fmt.Println(tree.find(23))
}