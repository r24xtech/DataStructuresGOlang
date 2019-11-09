package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  Value int
}

type Stack struct {
  nodes []*Node
  count int
}

// append adds elements onto the end of a slice.
func (s *Stack) Push(n *Node){
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}

func (s *Stack) Pop() {
  if s.count == 0 {
    defer func(){
      fmt.Println("Empty stack!")
      recover()
    }()
  } else {
    s.nodes = s.nodes[0:s.count]
    s.count--
  }
}

func add(s *Stack) {
  x := myutils.Get_int("Enter value: ")
  s.Push(&Node{x})
}

func pnt(s *Stack){
  fmt.Println("\ttop")
  for i := s.count - 1; i >= 0; i-- {
    fmt.Println(s.nodes[i].Value)
    //fmt.Println(*s.nodes[i])
  }
  fmt.Println("\tbottom")
}

func main() {
  s := new(Stack)

  fmt.Println("|| Welcome to the Stack ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: add(s)
    case 2: s.Pop()
    case 3: pnt(s)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
