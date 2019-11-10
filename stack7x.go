package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}

func (s *Stack) Push() {
  value := myutils.Get_int("Enter value: ")
  s.slice = append(s.slice, value)
}

func (s *Stack) Pop() {
  if s.Size() == 0 {
    defer func() {
      fmt.Println("The stack is empty!")
      recover()
    }()
  }
  s.slice = s.slice[0 : s.Size() - 1]
}

func (s *Stack) Size() int {
  return len(s.slice)
}

func sort(xs []int)  {
  var myslice []int = xs
  var firstNumber, secondNumber int
  var size int = len(xs)
  var t int = 1
  var count int

  myloop:for {
    count = 0
    for i := 0; i < size - t; i++ {
      firstNumber = myslice[i]
      secondNumber = myslice[i+1]
      if firstNumber > secondNumber {
        myslice[i] = secondNumber
        myslice[i+1] = firstNumber
        count++
      }
    }
    if count == 0 {
      break myloop
    }
    t++
  }
  fmt.Println(myslice)
}

func main() {
 s := new(Stack)

 fmt.Println("|| Welcome to the stack ||")
 fmt.Println("1-Add; 2-Delete; 3-Print; 4-Sort; 5-End \n\n")

 myloop:for {
   option := myutils.Get_int("Choose option: ")
   switch option {
    case 1: s.Push()
    case 2: s.Pop()
    case 3: fmt.Println(s.slice)
    case 4: sort(s.slice)
    case 5: break myloop
    default: fmt.Println("Invalid option.")
   }
 }
}
