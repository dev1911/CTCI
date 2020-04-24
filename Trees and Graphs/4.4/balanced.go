package main

import(
	"fmt"
	"math"
)

type Node struct{
	value int
	left *Node
	right *Node
	parent *Node
}

type BST struct{
	root *Node
} 

func (bst *BST)insert(value int){
	if bst.root == nil{
		bst.root = &Node{value:value , left:nil , right:nil , parent:nil}
	}else{
		bst.root.insert(value)
	} 
}

func (node *Node)insert(value int){
	if node == nil{
		return 
	}else if value < node.value{
		if node.left == nil{
			node.left = &Node{value:value , left:nil , right:nil , parent:node}
		}else{
			node.left.insert(value)
		}
	}else{
		if node.right == nil{
			node.right = &Node{value:value , left:nil , right:nil , parent:node}
		}else{
			node.right.insert(value)
		}
	}
}

func(bst *BST)inorder(){
	bst.root.inorder()
}

func(node *Node)inorder(){
	if node==nil{
		return
	}
	node.left.inorder()
	fmt.Printf(" %d " , node.value)
	node.right.inorder()
}

func(node *Node)findHeight()int{

	if node == nil{
		return -1
	}

	leftHeight := node.left.findHeight()
	if leftHeight == -2{
		return -2
	}

	rightHeight := node.right.findHeight()
	if rightHeight == -2{
		return -2
	}

	difference := math.Abs(float64(leftHeight - rightHeight))
	if difference > 1{
		return -2
	}else{
		return int(math.Max(float64(leftHeight) , float64(rightHeight))) + 1
	}
}

func(bst *BST)findBalanced()bool{
	return bst.root.findHeight() != -2
}

func main(){
	bst := &BST{}
	
	//Unbalanced tree
	// bst.insert(1)
	// bst.insert(2)
	// bst.insert(7)
	// bst.insert(5)
	// bst.insert(10)
	// bst.insert(8)

	//Balanced tree
	bst.insert(5)
	bst.insert(4)
	bst.insert(6)
	bst.insert(3)
	bst.insert(7)

	result := bst.findBalanced()
	if !result{
		fmt.Println("Tree is not balanced!")
	}else{
		fmt.Println("Tree is balanced!")
	}
}