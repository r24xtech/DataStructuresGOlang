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
  size int
  head *Node
  tail *Node
}

func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.head.previous = n
    n.next = l.head
    l.head = n
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
    keep := new(Node)
    keep = l.head
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
    l.head = keep
  }
}

func (l *List) RemoveFront() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty.")
  } else if l.size == 1 {
    l.head = l.head.next
    l.size--
  } else {
    l.head = l.head.next
    l.head.previous = nil
    l.size--
  }
}

func (l *List) RemoveBack() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty.")
  } else if l.size == 1 {
    l.tail = l.tail.previous
    l.size--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.size--
  }
}

/*
func (l *List) Swap() {
  l.head.next = l.head.next.next
  l.head.previous = l.head.next.previous
  l.head.previous.next = l.head
  l.head.next.previous = l.head
  l.head = l.head.previous
}

func (l *List) Swap() {
  keep := new(Node)
  keep = l.head.previous
  l.head.next = l.head.next.next
  l.head.previous = l.head.next.previous
  l.head.previous.next = l.head
  l.head.next.previous = l.head
  l.head.previous.previous = keep
  keep.next = l.head.previous
  l.head = l.head.previous
}
*/

func (l *List) Swap() {
  if l.size == 0 || l.size == 1 {
    fmt.Println("Can't swap!")
  } else if l.size == 2 {
    l.tail.next = l.head
    l.head.previous = l.tail
    l.head.next = nil
    l.tail.previous = nil
    l.tail = l.tail.next
    l.head = l.head.previous
  } else if l.size > 2 && l.head.previous != nil {
    keep := new(Node)
    keep = l.head.previous
    l.head.next = l.head.next.next
    l.head.previous = l.head.next.previous
    l.head.previous.next = l.head
    l.head.next.previous = l.head
    l.head.previous.previous = keep
    keep.next = l.head.previous
    l.head = l.head.previous
  } else {
    keep := new(Node)
    keep = l.head.next
    l.head.next = l.head.next.next
    l.head.next.previous = l.head
    l.head = keep
    l.head.next = keep.previous
    l.head.next.previous = l.head
    keep = nil // added
  }
}

func (l *List) Sort() {
  //keep := new(Node)
  //keep = l.head
  //l.head = l.head.next
  l.Swap()
  //l.head = keep

}

func main() {
  l := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print;")
  fmt.Println("[4] RemoveBack; [5] RemoveFront; [6] End;\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 0: l.Sort()
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
