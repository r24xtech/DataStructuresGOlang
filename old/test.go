package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) InsertFirst(i int) {
	data := &Node{data: i}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList) GetItemsFromStart() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
  l := new(LinkedList)

  l.InsertFirst(5)
  l.InsertFirst(9)
  l.InsertFirst(2)

  l.GetItemsFromStart()

  l.InsertFirst(13)
  l.InsertFirst(25)

  l.GetItemsFromStart()
}
x := myutils.Get_int("Enter value: ")
data := &Node{data: x}
s.size++
if s.head != nil {
	s.head.previous = data
	data.next = s.head
}
s.head = data


x := myutils.Get_int("Enter value: ")
data := &Node{data: x}
s.size++
if s.head == nil {
	s.head = data
	s.tail = data
	return
}
if s.tail != nil {
	s.tail.next = data
	data.previous = s.tail
}
s.tail = data
