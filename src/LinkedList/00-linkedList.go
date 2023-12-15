package main

import "fmt"

type LinkedList struct {
	head *Node
}

func (list *LinkedList) New() {
	list.head = nil
}
func (list *LinkedList) Add(data interface{}) {
	newNode := Node{data, nil}
	if list.head == nil {
		list.head = &newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
	}
}
func (list *LinkedList) Print() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

type Node struct {
	data interface{}
	next *Node
}

func main() {

	var list LinkedList
	list.New()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add("Hello")

	list.Print()
}
