# DataStructuresGOlang
Data structures in Go.

* If the struct has a pointer receiver on some of its methods, it is better to use it for the rest of the methods because it enables better consistency and predictability for the struct behaviors.
* `(s Stack)` vs. `(s *Stack)`: If you want to modify the data of a receiver from the method, the receiver must be a pointer.

```go
// spt.go
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
  size int
  current *Node
  last *Node
  keep *Node
}

func (q *Queue) Push() {
  x := myutils.Get_int("Enter value: ")
  qx := &Node{value: x, previous: nil}
  if q.size == 0 {
    q.current = qx
    q.last = q.current
  } else {
    q.last.previous = qx
    q.last = qx
  }
  q.size++
}

func (q *Queue) Pop() {
  if q.size == 0 {
    fmt.Println("Can't pop...the queue is empty!")
  } else {
    q.current = q.current.previous
    q.size--
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
    fmt.Println("\n")
    q.current = q.keep
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
```

<hr>

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

<hr>

```go
// y.go
// s.size not being used, but it might be useful
package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  next *Node
}

type Stack struct {
  size int
  top *Node
  bottom *Node
}

func (s *Stack) Push() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if s.top != nil {
    n.next = s.top
  }
  s.top = n
  s.size++
}

func (s *Stack) Pop() {
  if s.top == nil {
    fmt.Println("Can't pop...the stack is empty!")
  } else {
    s.top = s.top.next
    s.size--
  }
}

func (s *Stack) Print() {
  if s.top == nil {
    defer func() {
      fmt.Println("[...] Empty stack!")
      recover()
    }()
  } else {
    s.bottom = s.top
    fmt.Print("\ntop| ")
    for {
      fmt.Print(s.top.value, " ")
      if s.top.next == nil {
        break
      }
      s.top = s.top.next
    }
    fmt.Print(" |bottom\n\n")
    s.top = s.bottom
  }
}

func main() {
  s := new(Stack)

  fmt.Println("|| Welcome to the Stack ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: s.Push()
    case 2: s.Pop()
    case 3: s.Print()
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
```

<hr>

```go
// l.go
package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  previous *Node
  next *Node
}

type List struct {
  length int
  head *Node
  tail *Node
  keep *Node
}

func (l *List) InsertBack() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.length == 0 {
    l.head = n
    l.tail = l.head
  } else {
    formerTail := l.tail
    formerTail.next = n
    n.previous = formerTail
    l.tail = n
  }
  l.length++
}

func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.length == 0 {
    l.head = n
    l.tail = l.head
  } else {
    formerHead := l.head
    formerHead.previous = n
    n.next = formerHead
    l.head = n
  }
  l.length++
}

func (l *List) Print() {
  if l.length == 0 {
    fmt.Println("[...] Empty list!")
  } else {
    l.keep = l.head
    n := l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(n.value, " ")
      if n.next == nil {
        break
      }
      n = n.next
    }
    fmt.Print(" |back")
    l.head = l.keep
    fmt.Println("\n")
  }
}

func (l *List) RemoveFront() {
  if l.length == 0 {
    fmt.Println("Can't remove...empty list!")
  } else {
    l.head = l.head.next
    l.length--
  }
}

func (l *List) RemoveBack() {
  if l.length == 0 {
    fmt.Println("Can't remove...empty list!")
  } else if l.length == 1 {
    l.tail = l.tail.previous
    l.length--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.length--
  }
}

func main() {
  l := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print;")
  fmt.Println("[4] RemoveBack; [5] RemoveFront; [6] End;\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: l.InsertBack()
    case 2: l.InsertFront()
    case 3: l.Print()
    case 4: l.RemoveBack()
    case 5: l.RemoveFront()
    case 6: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}

``

<hr>

```go
// dll.go
package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  next *Node // next moves to the back of the list
  previous *Node // previous moves to the front of the list
}

type List struct {
  size int
  head *Node
  tail *Node
}

func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.head.previous = n
    n.next = l.head
    l.head = n
  }
  l.size++
}

func (l *List) InsertBack() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.next = n
    n.previous = l.tail
    l.tail = n
  }
  l.size++
}

func (l *List) Print() {
  if l.size == 0 {
    fmt.Println("[...] the list is empty!")
  } else {
    keep := new(Node)
    keep = l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(l.head.value, " ")
      if l.head.next == nil {
        break
      }
      l.head = l.head.next
    }
    fmt.Print(" |back")
    fmt.Println("\n")
    l.head = keep
  }
}

func (l *List) RemoveBack() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.tail = l.tail.previous
    l.size--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.size--
  }
}

func (l *List) RemoveFront() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.head = l.head.next
    l.size--
  } else {
    l.head = l.head.next
    l.head.previous = nil
    l.size--
  }
}

func (l *List) RemoveAt() {
  x := myutils.Get_int("Remove by index: ")
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    if x == 0 {
      l.RemoveFront()
    } else {
      fmt.Println("Invalid index.")
    }
  } else if l.size == 2 {
    if x == 0 {
      l.RemoveFront()
    } else if x == 1 {
      l.RemoveBack()
    } else {
      fmt.Println("Invalid index.")
    }
  } else {
    if x == 0 {
      l.RemoveFront()
    } else if x == l.size-1 {
      l.RemoveBack()
    } else {
      if x > 0 && x < l.size-1 {
        keep := new(Node)
        keep = l.head
        for i := 0; i < x; i++ {
          l.head = l.head.next
        }
        l.head.previous.next = l.head.next
        l.head.next.previous = l.head.previous
        l.size--
        //l.RemoveFront()
        l.head = keep
      } else {
        fmt.Println("Invalid index.")
      }
    }
  }
}

func (l *List) InsertAt() {
  x := myutils.Get_int("Insert by index: ")
  if l.size == 0 {
    l.InsertFront()
  } else if l.size == 1 {
    if x == 0 {
      l.InsertFront()
    } else if x == 1 {
      l.InsertBack()
    } else {
      fmt.Println("Invalid index.")
    }
  } else {
    if x == 0 {
      l.InsertFront()
    } else if x == l.size {
      l.InsertBack()
    } else if x > 0 && x < l.size {
      n := new(Node)
      keep := new(Node)
      keep = l.head
      for i := 0; i < x; i++ {
        l.head = l.head.next
      }
      z := myutils.Get_int("Enter value: ")
      n.value = z
      n.next = l.head
      n.previous = l.head.previous
      l.head.previous = n
      l.head.previous.previous.next = n
      l.size++
      //l.InsertFront()
      l.head = keep
    } else {
      fmt.Println("Invalid index.")
    }
  }
}

func main() {
  l := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print; [4]InsertByIndex;")
  fmt.Println("[5] RemoveBack; [6] RemoveFront; [7]RemoveByIndex; [8] End;\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: l.InsertBack()
    case 2: l.InsertFront()
    case 3: l.Print()
    case 4: l.InsertAt()
    case 5: l.RemoveBack()
    case 6: l.RemoveFront()
    case 7: l.RemoveAt()
    case 8: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}

```
