# DataStructuresGOlang
Data structures in Go.

* If the struct has a pointer receiver on some of its methods, it is better to use it for the rest of the methods because it enables better consistency and predictability for the struct behaviors.
* `(s Stack)` vs. `(s *Stack)`: If you want to modify the data of a receiver from the method, the receiver must be a pointer.

```go
// stack3.go
package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}

func (s *Stack) Push(i int) {
  s.slice = append(s.slice, i)
}

func (s *Stack) Peek() int {
  return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() {
  if s.Size() == 0 {
    defer func() {
      fmt.Println("Empty stack!")
      recover()
    }()
  }
  s.slice = s.slice[0 : len(s.slice)-1]
}

func (s *Stack) Size() int {
  return len(s.slice)
}

func add(s *Stack) {
  x := myutils.Get_int("Enter integer: ")
  s.Push(x)
}

func main() {
  s := new(Stack)

  fmt.Println("|| Welcome to the Stack ||")
  fmt.Println("1- Add; 2-Delete; 3-Print; 4-End\n\n")

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
```

<hr>

```go
// queue.go
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
```
