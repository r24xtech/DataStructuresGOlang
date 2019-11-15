package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type List struct {
  size int
  start *Node
}

func (l *List) Push(newNode *Node) {
  if l.size == 0 {
    l.start = newNode
  } else {
    currentNode := l.start
    for currentNode.next != nil {
      currentNode = currentNode.next
    }
    currentNode.next = newNode
  }
  l.size++
}

func print(l *List) {
  for {
    fmt.Println(l.start.value)
    if l.start.next == nil {
      break
    }
    l.start = l.start.next
  }
}

func main() {
  l := new(List)

  l.Push(&Node{value: 5})
  l.Push(&Node{value: 9})

  print(l)
  l.Push(&Node{value: 3})
  print(l)
}
