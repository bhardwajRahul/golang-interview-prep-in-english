package main

func main() {

	// make is only for slices, maps, and channels;
	// it initializes their internal runtime data structures and returns the value itself, not a pointer.
	// Use make when creating usable slices, maps, or channels.
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	println("Map value for 'one':", m["one"])
	println("Map value for 'two':", m["two"])

	// new(T) allocates zeroed memory for type T and returns a pointer: *T
	// Use new when you want a pointer to a zero value.
	s := new(int)
	*s = 42
	println("Value of s:", *s)
}
