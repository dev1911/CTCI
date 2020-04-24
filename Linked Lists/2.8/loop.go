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

func(l *LinkedList)detectLoop() *Node{
	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil{
		fast = fast.next.next
		slow = slow.next

		if fast == slow{
			break
		}
	}

	if fast == nil || fast.next == nil{
		return nil
	}

	slow = l.head
	for slow != fast{
		slow = slow.next
		fast = fast.next
	}

	return fast
}

func main(){
	l1 := LinkedList{}
	l1.addNodeBeginning(1)
	ptr := l1.head
	l1.addNodeBeginning(2)
	l1.addNodeBeginning(3)
	l1.addNodeBeginning(4)
	ptr2 := l1.head
	l1.addNodeBeginning(5)
	l1.addNodeBeginning(6)
	l1.addNodeBeginning(7)

	ptr.next = ptr2

	result := l1.detectLoop()

	fmt.Printf("Loop detected!\nLoop start found at %d" , result.value)	
}
/*
	7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
				   |______________|
*/