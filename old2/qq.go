package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  previous *Node
}

type Queue struct {
  current *Node
  last *Node
  keep *Node
  size uint64
}

func (q *Queue) Push() {
  x := myutils.Get_int("Enter value: ")
  if q.size == 0 {
    q.current = &Node{value: x, previous: nil}
    q.last = q.current
    q.size++
  } else {
    qx := &Node{value: x, previous: nil}
    q.last.previous = qx
    q.last = qx
    q.size++
  }
}

func (q *Queue) Pop() {
  if q.size > 0 {
    q.current = q.current.previous
    q.size--
  } else {
    fmt.Println("Can't pop...the queue is empty!")
  }
}

func (q *Queue) Print() {
  if q.size == 0 {
    fmt.Println("[...] The queue is empty!")
  } else {
    q.keep = q.current
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(q.current.value, " ")
      if q.current.previous == nil {
        break
      }
      q.current = q.current.previous
    }
    fmt.Print(" |back")
    q.current = q.keep
    fmt.Println("\n")
  }
}

func main() {
  q := new(Queue)

  fmt.Println("|| Welcome to the Queue ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: q.Push()
    case 2: q.Pop()
    case 3: q.Print()
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
