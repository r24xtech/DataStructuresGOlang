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
  size int
  top *Node
}

func (s *Stack) Push(x int) {
  n := new(Node)
  n.value =x
  if s.top != nil {
    n.next = s.top
  }
  s.top = n
  s.size++
}

func (s *Stack) Pop() (int, bool){
  if s.top == nil {
    fmt.Println("The stack is empty!")
    return 0, false
  } else {
    i := s.top.value
    s.top = s.top.next
    s.size--
    return i, true
  }
}

func main() {
  s := new(Stack)

  t := myutils.Get_int("How many items: ")
  for i := 0; i < t; i++ {
    value := myutils.Get_int("Enter value: ")
    s.Push(value)
  }

  for {
    a, b := s.Pop()
    if b == false {
      break
    }
    fmt.Println(a)
  }
}
