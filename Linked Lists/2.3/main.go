package main

import(
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addNodeBeginning(value int) {
	newNode := &Node{
		value: value,
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) displayLinkedList() {
	temp := l.head
	for temp != nil {
		fmt.Printf(" %d ", temp.value)
		temp = temp.next
	}
	fmt.Println("")
}

func (l *LinkedList)getPointerTo(value int) *Node{
	temp := l.head
	for temp != nil{
		if temp.value == value{
			return temp
		}
		temp = temp.next
	}
	return nil
}

func(node *Node)delete() bool{
	if node == nil || node.next == nil{
		return false
	}
	
	node.value = node.next.value
	node.next = node.next.next
	return true
}

func main(){
	l1 := LinkedList{}
	l1.addNodeBeginning(1)
	l1.addNodeBeginning(2)
	l1.addNodeBeginning(3)
	l1.addNodeBeginning(4)

	l1.displayLinkedList()
	ptr := l1.getPointerTo(3)
	result := ptr.delete()
	if result{
		fmt.Println("Delete successful")
		l1.displayLinkedList()
	}else{
		fmt.Println("Delete failed")
	}
}