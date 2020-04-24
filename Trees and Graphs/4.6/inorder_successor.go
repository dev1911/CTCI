package main

import(
	"fmt"
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

func(bst *BST)findNextNode(node *Node)int{
	//If right subtree exists, find minimum value in right subtree
	//Else follow parent link until node whose right subtree is not explored
	if node.right != nil{
		temp1 := node.right

		for temp1!=nil && temp1.left!=nil{
			temp1 = temp1.left
		}
		return temp1.value
	}else if node.right == nil{
		temp1 := node
		temp2 := temp1.parent
		for temp2!= nil && temp2.left != temp1{
			temp1 = temp2
			temp2 = temp2.parent
		}
		return temp2.value
	}

	return -1
}

func(bst *BST)findNode(value int)*Node{
	return bst.root.findNode(value)
}

func(root *Node)findNode(value int)*Node{

	if root == nil{
		return nil
	}
	if value < root.value{
		return root.left.findNode(value)
	}else if value > root.value{
		return root.right.findNode(value)
	}else if value == root.value{
		return root
	}

	return nil
}

func(bst *BST)findInorderSuccessor(value int)int{
	ptr := bst.findNode(value)
	return bst.findNextNode(ptr)
}

//Driver program
func main(){
	bst := &BST{}
	bst.insert(1)
	bst.insert(2)
	bst.insert(7)
	bst.insert(5)
	bst.insert(10)
	bst.insert(8)

	fmt.Println("The inorder traversal of the BST is ")
	bst.inorder()
	result := bst.findInorderSuccessor(8)
	fmt.Printf("\nThe inorder successor of 8 is %d",result)
}

/*
		1
		   2
		      7 
		     5  10
		       8 
*/