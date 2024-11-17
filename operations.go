package main

// Basic stack operations
func (s *Stack) Swap() bool {
	if s.Size < 2 {
		return false
	}

	first := s.Head.Value
	second := s.Head.Next.Value
	s.Head.Value = second
	s.Head.Next.Value = first
	return true
}

func Push(from, to *Stack) bool {
	if from.Size == 0 {
		return false
	}

	value, _ := from.Pop()
	to.Push(value)
	return true
}

func (s *Stack) Rotate() bool {
	if s.Size <= 1 {
		return false
	}

	value := s.Head.Value
	s.Pop()

	current := s.Tail
	newNode := &Node{Value: value, Prev: current}
	current.Next = newNode
	s.Tail = newNode
	s.Size++
	return true
}

func (s *Stack) ReverseRotate() bool {
	if s.Size <= 1 {
		return false
	}

	value := s.Tail.Value
	s.Tail = s.Tail.Prev
	s.Tail.Next = nil
	s.Push(value)
	return true
}
