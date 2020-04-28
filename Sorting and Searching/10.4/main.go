package main

import(
	"fmt"
)

type Listy struct{
	arr []int
}

func(l *Listy)insert(value int){
	l.arr = append(l.arr , value)
}

func(l *Listy)getElement(i int) int{
	if i > len(l.arr){
		return -1
	}else{
		return l.arr[i]
	}
}

func(l *Listy)search(value int)int{
	idx := 1
	for l.getElement(idx) != -1 && l.getElement(idx) < value{
		idx = idx * 2
	}
	return l.binarySearch(value, 0 , idx)
}

func(l *Listy)binarySearch(value int , low int , high int)int{
	var mid int
	var middle int
	for low <= high{
		mid = (low + high)/2
		middle = l.getElement(mid)

		if middle > value || middle == -1{
			high = mid - 1
		}else if middle < value{
			low = mid + 1
		}else{
			return mid
		}
	}
	return -1
}

func main(){
	list := &Listy{}
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.insert(8)
	list.insert(9)

	result := list.search(9)
	fmt.Printf("Element found at position %d" , result)
}