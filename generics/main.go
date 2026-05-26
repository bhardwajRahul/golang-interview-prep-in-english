package main

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false // Return zero value and false if stack is empty
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1] // Remove the last element
	return element, true
}

func main() {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)

	if element, ok := stack.Pop(); ok {
		println("Popped element:", element)
	} else {
		println("Stack is empty")
	}
}
