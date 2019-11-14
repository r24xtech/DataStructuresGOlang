package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
	data int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.data = x
  if t.root == nil {
    t.root = n
  } else {
    parent := t.root
    myloop:for { // you need the for loop
      switch {
      case x < parent.data:
        if parent.left == nil {
          parent.left = n
          break myloop
        }
        parent = parent.left
      case x > parent.data:
        if parent.right == nil {
          parent.right = n
          break myloop
        }
        parent = parent.right
      default:
        break myloop
      }
    }
  }
}

func preOrder(n *Node) {
  if n == nil {
    return
  }
  fmt.Println(n.data)
  preOrder(n.left)
  preOrder(n.right)
}

func main() {
  t := new(Tree)

  t.Insert() // 2
  t.Insert() // 1
  t.Insert() // 3
  t.Insert() // 9
  t.Insert() // 8

  preOrder(t.root)

  //fmt.Println(t.root.data)
  //fmt.Println(t.root.right.data)
  //fmt.Println(t.root.left.data)
  //fmt.Println(t.root.right.right.data)
  //fmt.Println(t.root.right.right.left.data)

}
