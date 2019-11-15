package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}
// Push adds the integer provided to the top of the stack
func (s *Stack) Push(i int) {
  s.slice = append(s.slice, i)
}

// Peek returns the top item from the stack, but DOES NOT remove it from the stack
func (s *Stack) Peek() int {
  return s.slice[len(s.slice)-1]
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() int {
  if s.Size() == 0 {
    fmt.Println("Empty stack!")
    return 0
  }
  var ret int = s.Peek()
  s.slice = s.slice[0 : len(s.slice)-1]
  return ret
}

func (s *Stack) Size() int {
  return len(s.slice)
}

func (s *Stack) String() string {
  return fmt.Sprint(s.slice)
}

func add(s *Stack) {
  value := myutils.Get_int("Enter value: ")
  s.Push(value)
}

func main() {
  var s *Stack = new(Stack)

  fmt.Println("|| Welcome to the Stack (LIFO)||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: add(s)
    case 2: s.Pop()
    case 3: fmt.Println(s)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
