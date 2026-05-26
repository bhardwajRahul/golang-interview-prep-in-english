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

type StackInterface[T any] interface {
	Push(element T)
	Pop() (T, bool)
}

type StackWrapper[T any] struct {
	StackInterface[T] // Embedding the generic Stack with int type
}

func (s *StackWrapper[T]) Push(element any) {
	s.StackInterface.Push(element.(T))
}

func (s *StackWrapper[T]) Pop() (any, bool) {
	element, ok := s.StackInterface.Pop()
	return element, ok
}

func main() {
	stack := StackWrapper[int]{StackInterface: &Stack[int]{}}
	stack.Push(10)
	stack.Push(20)

	if element, ok := stack.Pop(); ok {
		println("Popped element:", element.(int))
	} else {
		println("Stack is empty")
	}
}
