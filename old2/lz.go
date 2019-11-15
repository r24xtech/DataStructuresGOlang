package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  next *Node
  previous *Node
}

type List struct {
  size uint64
  head *Node
  tail *Node
  keep *Node
}
/*
next goes to the back ==>
FRONT| x y z a b c d e f |BACK
<== previous goes to the front
*/
func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.head.previous = n // sets new node in front of the list
    n.next = l.head // makes connection between "new node" and "old head"
    l.head = n // says the new head is the new node
  }
  l.size++
}

func (l *List) InsertBack() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.next = n
    n.previous = l.tail
    l.tail = n
  }
  l.size++
}

func (l *List) Print() {
  if l.size == 0 {
    fmt.Println("[...] the list is empty!")
  } else {
    l.keep = l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(l.head.value, " ")
      if l.head.next == nil {
        break
      }
      l.head = l.head.next
    }
    fmt.Print(" |back")
    fmt.Println("\n")
    l.head = l.keep
  }
}

func (l *List) RemoveBack() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.tail = l.tail.previous
    l.size--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.size--
  }
}

func (l *List) RemoveFront() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.head = l.head.next
    l.size--
  } else {
    l.head = l.head.next
    l.head.previous = nil
    l.size--
  }
}

func main() {
  l := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print;")
  fmt.Println("[4] RemoveBack; [5] RemoveFront; [6] End;\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: l.InsertBack()
    case 2: l.InsertFront()
    case 3: l.Print()
    case 4: l.RemoveBack()
    case 5: l.RemoveFront()
    case 6: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
