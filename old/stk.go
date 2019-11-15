package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  previous *Node
}

type Stack struct {
  size int
  top *Node
}

func (s *Stack) Push() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  n.previous = s.top
  s.top = n
  s.size++
}

func (s *Stack) Pop(n *Node) {
  if s.size > 0 {
    n = s.top
    s.top = s.top.previous
    s.size--
  } else {
    fmt.Println("The stack is empty!")
  }
}

func print(s *Stack) {
  fmt.Println()
  if s.size > 0 {
    for {
      fmt.Print( s.top.value, "  ")
      if s.top.previous == nil {
        break
      }
      s.top = s.top.previous
    }
  } else {
    fmt.Println("...")
  }
}

func main() {
  s := new(Stack)

  fmt.Println("Welcome to the xStack")
  fmt.Println("1-Add; 2-Delte; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("\nChoose option: ")
    switch option {
    case 1: s.Push()
    case 2: s.Pop(&Node{})
    case 3: print(s)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
