package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  previous *Node
  next *Node
}

type List struct {
  length int
  head *Node
  tail *Node
  keep *Node
}

func (l *List) InsertBack() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.length == 0 {
    l.head = n
    l.tail = l.head
  } else {
    formerTail := l.tail
    formerTail.next = n
    n.previous = formerTail
    l.tail = n
  }
  l.length++
}

func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.length == 0 {
    l.head = n
    l.tail = l.head
  } else {
    formerHead := l.head
    formerHead.previous = n
    n.next = formerHead
    l.head = n
  }
  l.length++
}

func (l *List) Print() {
  if l.length == 0 {
    fmt.Println("[...] Empty list!")
  } else {
    l.keep = l.head
    n := l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(n.value, " ")
      if n.next == nil {
        break
      }
      n = n.next
    }
    fmt.Print(" |back")
    l.head = l.keep
    fmt.Println("\n")
  }
}

func (l *List) RemoveFront() {
  if l.length == 0 {
    fmt.Println("Can't remove...empty list!")
  } else {
    l.head = l.head.next
    l.length--
  }
}

func (l *List) RemoveBack() {
  if l.length == 0 {
    fmt.Println("Can't remove...empty list!")
  } else if l.length == 1 {
    l.tail = l.tail.previous
    l.length--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.length--
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
