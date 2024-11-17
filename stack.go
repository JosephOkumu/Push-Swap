package main

type Node struct {
    Value int
    Next  *Node
    Prev  *Node
}

type Stack struct {
    Head   *Node
    Tail   *Node
    Size   int
    Name   string
}

func NewStack(name string) *Stack {
    return &Stack{
        Name: name,
        Size: 0,
    }
}

func (s *Stack) Push(value int) {
    newNode := &Node{Value: value}
    if s.Head == nil {
        s.Head = newNode
        s.Tail = newNode
    } else {
        newNode.Next = s.Head
        s.Head.Prev = newNode
        s.Head = newNode
    }
    s.Size++
}

func (s *Stack) Pop() (int, bool) {
    if s.Head == nil {
        return 0, false
    }
    value := s.Head.Value
    s.Head = s.Head.Next
    if s.Head != nil {
        s.Head.Prev = nil
    } else {
        s.Tail = nil
    }
    s.Size--
    return value, true
}

func (s *Stack) IsSorted() bool {
    if s.Size <= 1 {
        return true
    }
    current := s.Head
    for current.Next != nil {
        if current.Value > current.Next.Value {
            return false
        }
        current = current.Next
    }
    return true
}