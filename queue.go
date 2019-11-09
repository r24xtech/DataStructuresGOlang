package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Queue struct {
  slice []int // slices can be resized dynamically
}
// index = 0 is the front of the queue

// adds the integer provided to the back of the queue
func (q * Queue) Enqueue(i int) {
  q.slice = append(q.slice, i)
}

func (q *Queue) Dequeue() {
  if q.Size() == 0 {
    defer func(){
      fmt.Println("Empty queue!")
      recover()
    }()
  }
  q.slice = q.slice[1 : len(q.slice)]
}

func (q *Queue) Size() int {
  return len(q.slice)
}

func Add(q *Queue) {
  value := myutils.Get_int("Enter value: ")
  q.Enqueue(value)
}

func main() {
  var q *Queue = new(Queue)

  fmt.Println("|| Welcome to the Queue (FIFO)||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End \n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: Add(q)
    case 2: q.Dequeue()
    case 3: fmt.Println("front| ",q.slice, "|back")
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
