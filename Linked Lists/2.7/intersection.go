package main

import(
	"fmt"
)

type Node struct{
	value int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList)addNodeBeginning(value int){
	newNode := &Node{
		value : value,
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList)displayLinkedList(){
	temp := l.head
	for temp != nil{
		fmt.Printf(" %d ",temp.value)
		temp = temp.next
	}
	fmt.Println("")
}


func findIntersection(head1 *Node , head2 *Node) *Node{
	
	//Finding out length of both lists
	temp1 := head1
	var length1 int
	for temp1 != nil{
		temp1 = temp1.next
		length1++
	}

	temp2 := head2
	var length2 int
	for temp2 != nil{
		temp2 = temp2.next
		length2++
	}

	//If they dont have common end node, they do not intersect
	if temp1 != temp2{
		return nil
	}

	//Advancing pointer to longer list by the difference of the lengths
	temp1 = head1
	temp2 = head2

	if(length1 > length2){
		for i :=0 ; i< length1 - length2 ; i++{
			temp1 = temp1.next
		}
	}else{
		for i := 0 ; i<length2 - length1 ; i++{
			temp2 = temp2.next
		}
	}

	//Advancing one node at a time till both pointers are same
	for temp1 != temp2{
		temp1 = temp1.next
		temp2 = temp2.next
	}

	return temp1
}

func main(){
	l1 := LinkedList{}
	l1.addNodeBeginning(1)
	l1.addNodeBeginning(2)
	l1.addNodeBeginning(3)
	l2 := LinkedList{head:l1.head}
	l1.addNodeBeginning(4)
	l1.addNodeBeginning(5)
	l2.addNodeBeginning(6)
	l2.addNodeBeginning(7)
	l2.addNodeBeginning(8)
	l2.addNodeBeginning(9)

	result := findIntersection(l1.head , l2.head)
	if result != nil{
		fmt.Printf("Intersection found at %d" , result.value)
	}else{
		fmt.Printf("No intersection of lists")
	}
}

/*
	l1
   	 5 4
   	 	 3 2 1
 9 8 7 6
l2
*/