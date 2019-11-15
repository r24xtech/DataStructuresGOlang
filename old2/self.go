package main

import (
  "fmt"
  //"github.com/r24xtech/myutils"
)
/*
func Insert(l []int) {
  x := myutils.Get_int("Enter value: ")

  if len(l) == 0 {
    l = append(l, x)
  } else {
    for i := 0; i < len(l); i++ {

    }
  }
}*/

func InsertAt(l []int, i int, x int) []int {
  var l1 []int
  var l2 []int
  var lx []int

  l1 = l[:i]
  l2 = l[i:]
  fmt.Println(l1,l2)
  copy(lx, l2)

  fmt.Println(lx)

  //l1 = append(l1,x)
  //fmt.Println(l1,l2)
  return l1
}

func main() {
  //var list [] int

  slice := []int{1,2,3,4,5}
  op:=InsertAt(slice, 2, 9)
  fmt.Println(op)
  /*myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: Insert(list)
    case 2: break myloop
    default: fmt.Println("Invalid option.")
    }
  }*/
}
