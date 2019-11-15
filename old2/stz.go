package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  next *Node
}

type Stack struct {
  size uint64
  top *Node
  keep *Node
}

func (s *Stack) Push() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if s.size != 0 {
    n.next = s.top
  }
  s.top = n
  s.size++
}

func (s *Stack) Pop() {
  if s.size == 0 {
    fmt.Println("Can't pop...the stack is empty!")
  } else {
    s.top = s.top.next
    s.size--
  }
}

func (s *Stack) Print() {
  if s.size == 0 {
    fmt.Println("[...] the stack is empty!")
  } else {
    s.keep = s.top
    fmt.Println()
    fmt.Print("top| ")
    for {
      fmt.Print(s.top.value, " ")
      if s.top.next == nil {
        break
      }
      s.top = s.top.next
    }
    fmt.Print(" |bottom")
    fmt.Println("\n")
    s.top = s.keep
  }
}

func main() {
  s := new(Stack)

  fmt.Println("|| Welcome to the Stack ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: s.Push()
    case 2: s.Pop()
    case 3: s.Print()
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
