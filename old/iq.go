package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}

type Queue struct {
  slice []int
}

type dataStructures interface {
  Push() []int
  Pop() []int
  Size() int
}

func (s *Stack) Push() []int {
  x := myutils.Get_int("Enter value: ")
  s.slice = append(s.slice, x)
  return s.slice
}

func (q *Queue) Push() []int {
  x := myutils.Get_int("Enter value: ")
  q.slice = append(q.slice, x)
  return q.slice
}

func (s *Stack) Pop() []int {
  if s.Size() == 0 {
    defer func() {
      fmt.Println("The stack is empty!")
      recover()
    }()
  }
  s.slice = s.slice[0 : s.Size() - 1]
  return s.slice
}

func (q *Queue) Pop() []int {
  if q.Size() == 0 {
    defer func() {
      fmt.Println("The queue is empty!")
      recover()
    }()
  }
  q.slice = q.slice[1 : q.Size()]
  return q.slice
}

func (s *Stack) Size() int {
  return len(s.slice)
}

func (q *Queue) Size() int {
  return len(q.slice)
}

func pushDs(ds ...dataStructures) [][]int {
  //var allDs [][]int
  allDs := make([][]int,len(ds)-2)
  for _, v := range ds {
    allDs = append(allDs, v.Push())
  }
  return allDs
}

func popDs(ds ...dataStructures) [][]int {
  allDs := make([][]int, len(ds)-2)
  for _, v := range ds {
    allDs = append(allDs, v.Pop())
  }
  return allDs
}

func Print(a [][]int) {
  if len(a) > 0 {
    fmt.Println("stack: ", a[0], "\nqueue: ", a[1])
  } else {
    fmt.Println("stack: []\nqueue: []")
  }
}

func main() {
  s := new(Stack)
  q := new(Queue)

  var a[][]int
  fmt.Println("1-Push; 2-Pop; 3-Print; 4-End\n\n")
  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: a = pushDs(s, q)
    case 2: a = popDs(s, q)
    case 3: Print(a)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
