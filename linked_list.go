package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type List struct {
  head *Node
}

func (l *List) addBack (num int) {
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

func (l *List) addFront (num int) {
  new_data := &Node{value: num}
  if l.head != nil {
    new_data.next = l.head
    l.head = new_data
  } else {
    l.head = new_data
  }
}

func (l *List) popBack () {
  current := l.head
  for current.next.next != nil {
    current = current.next
  }
  current.next = nil
}

func (l *List) popFront () {
  if l.head != nil {
    l.head = l.head.next
  }
}

func (l *List) print() ([]int) {
  result := []int {}
  current := l.head
  if current != nil {
    result = append(result, current.value)
    // fmt.Print(current.value, "\n")
  }
  for current.next != nil {
    result = append(result, current.next.value)
    // fmt.Print(current.next.value, " ")
    current = current.next
  }

  return result
}
