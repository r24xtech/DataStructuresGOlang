package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  data int
  next *Node
  previous *Node
}

type Stack struct {
  head *Node
  tail *Node
  size int
}

func (s *Stack) Push() {
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
}

func (s *Stack) Pop() {
  if s.size == 0 {
    fmt.Println("The stack is empty!")
  } else {
    fmt.Println(s.head)
    fmt.Println(s.tail)
    s.tail = nil
    s.tail.next = s.tail
    s.size--
  }
}

func (s *Stack) Print() {
  current := s.head
  fmt.Print("bottom| ")
  for current != nil {
    fmt.Print(current, " ")
    current = current.next
  }
  fmt.Print(" |top")
  fmt.Println()
}

func main() {
  s := new(Stack)

  fmt.Println("|| Welcome to the Stack (LIFO)||")
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
