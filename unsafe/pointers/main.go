package main

import "unsafe"

type MyStruct struct {
	value []unsafe.Pointer
}

func (s *MyStruct) Reset() {
	for i := range s.value {
		s.value[i] = nil
	}
}

func main() {
	s := MyStruct{
		value: make([]unsafe.Pointer, 10),
	}

	for i := 0; i < 10; i++ {
		s.value[i] = unsafe.Pointer(&i)
	}
	for i := 0; i < 10; i++ {
		println(*(*int)(s.value[i]))
	}
	s.Reset()
	for i := 0; i < 10; i++ {
		println((s.value[i]))
	}
}
