package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}

func (s *Stack) Push(x int) {
  s.slice = append(s.slice, x)
}

func (s *Stack) Pop() {
  if s.Size() == 0 {
    defer func(){
      fmt.Println("The stack is empty!")
      recover()
    }()
  }
  s.slice = s.slice[0 : s.Size() - 1]
}

func (s *Stack) Size() int{
  return len(s.slice)
}

func add(s *Stack) {
  value := myutils.Get_int("Enter value: ")
  s.Push(value)
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
    case 3: fmt.Println(s.slice)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
