package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type Stack struct {
  size int
  start *Node
}

func (s *Stack) Push(newNode *Node) {
  if s.size == 0 {
    s.start = newNode
  } else {
    currentNode := s.start
    for currentNode.next != nil {
      currentNode = currentNode.next
    }
    currentNode.next = newNode
  }
  s.size++
}

func main() {
  s := new(Stack)

  s.Push(&Node{value: 5})
  s.Push(&Node{value: 9})

  for {
    fmt.Println(s.start.value)
    if s.start.next == nil {
      break
    }
    s.start = s.start.next
  }
}
