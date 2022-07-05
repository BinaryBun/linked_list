package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type List struct {
  head *Node
}

func (l *List) addBack(num int) {
  new_data := &Node{value: num}
  if l.head == nil {
    l.head = new_data
  } else {
    current := l.head
    for current.next != nil {
      current = current.next
    }
    current.next = new_data
  }
}

func (l *List) print() {
  current := l.head
  if current != nil {
    fmt.Print(current.value, " ")
  }
  for current.next != nil {
    fmt.Print(current.next.value, " ")
    current = current.next
  }
}

func main() {
  data := List{}
  data.addBack(12)
  data.addBack(14)
  data.addBack(5)
  data.addBack(20)
  data.print()
}
