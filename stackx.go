package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)
/* using interface field unnecessarily, you'll only add integters */
type Node struct {
  item interface{}
  next *Node
}

type Stack struct {
  sp *Node
  size uint64
}

func New() *Stack {
  var stack *Stack = new(Stack)
  stack.size = 0
  return stack
}

func (s *Stack) Push(item interface{}) {
  s.sp = &Node{item: item, next: s.sp}
  s.size++
}

func (s *Stack) Pop() interface{}{
  if s.size > 0 {
    i := s.sp.item
    s.sp = s.sp.next
    s.size--
    return i
  }
  return nil
}

func print(s *Stack) {
  for s.size > 0{
    fmt.Println(s.Pop())
  }
}

func add(s *Stack) {
  x := myutils.Get_int("Enter value: ")
  s.Push(x)
}

func main() {
  s := New()

  fmt.Println("Problem: you lose memory after pop...")
  fmt.Println("3-You also lose all items from stack")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
  option := myutils.Get_int("Choose option: ")
  switch option {
  case 1: add(s)
  case 2: s.Pop()
  case 3: print(s)
  case 4: break myloop
  default: fmt.Println("Invalid option.")
  }
}
}
