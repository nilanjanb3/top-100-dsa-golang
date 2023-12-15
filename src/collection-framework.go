package main

import (
	"container/list"
	"fmt"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/golang-collections/collections/stack"
)

func main() {
	fmt.Println("~~~~~~~~~~~~~~~Doubly Linked List~~~~~~~~~~~~~~~")
	myList := list.New()

	myList.PushFront(20)
	myList.PushFront(10)
	myList.PushBack(30)
	fmt.Println("Left --> Right")
	i := myList.Front()
	for i != nil {
		fmt.Println(i.Value)
		i = i.Next()
	}
	myList.Remove(myList.Back())
	fmt.Println("Right --> Left")
	i = myList.Back()
	for i != nil {
		fmt.Println(i.Value)
		i = i.Prev()
	}
	fmt.Println("~~~~~~~~~~~~~~~Stack~~~~~~~~~~~~~~~")
	var stack stack.Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push("Hello")

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
	fmt.Println("~~~~~~~~~~~~~~~Queue~~~~~~~~~~~~~~~")
	var queue queue.Queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	for queue.Len() > 0 {
		fmt.Println(queue.Dequeue())
	}
	fmt.Println("~~~~~~~~~~~~~~~Set~~~~~~~~~~~~~~~")
	set := set.New()
	set.Insert(10)
	set.Insert(20)
	set.Insert(30)

	fmt.Println(set.Has(20))
	fmt.Println(set.Has(40))

	fmt.Println("~~~~~~~~~~~~~~~Map~~~~~~~~~~~~~~~")

	myMap := make(map[string]int)
	text := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(text, " ")
	for _, word := range words {
		myMap[word]++
		// fmt.Println(word)
	}
	fmt.Println(myMap)

	myPQ := pq.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	fmt.Println("~~~~~~~~~~~~~~~Priority Queue~~~~~~~~~~~~~~~")
	myPQ.Enqueue(50)
	myPQ.Enqueue(70)
	myPQ.Enqueue(10)

	for myPQ.Size() != 0 {
		fmt.Println(myPQ.Dequeue())
	}
}
