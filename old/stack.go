package main
/* LIFO */
import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  Value int
}

func (n *Node) String() string {
  return fmt.Sprint(n.Value)
}

//NewStack returns a new stack
func NewStack() *Stack {
  return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed
type Stack struct {
  nodes []*Node
  count int
}

// Push adds a node to the stack
func (s *Stack) Push(n *Node) {
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
  if s.count == 0 {
    return nil
  }
  s.count--
  return s.nodes[s.count]
}

func add(s *Stack) {
  value := myutils.Get_int("Enter value: ")
  s.Push(&Node{value})
}

func pnt(s *Stack) {
  fmt.Println(s.nodes)
}

func main() {
  s := NewStack()

  fmt.Println("|| Welcome to the Stack (LIFO)||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: add(s)
    case 2: s.Pop()
    case 3: pnt(s)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
