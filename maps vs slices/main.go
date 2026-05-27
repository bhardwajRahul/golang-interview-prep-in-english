package main

func main() {

	//In Go, slices and maps are both reference-like data structures, but they solve different problems.

	// hmap for the map header and bmap for buckets.

	//Maps:

	// A map is a small header:
	// type hmap struct {
	//	count     int
	//flags     uint8
	//B         uint8  // log_2 of the number of buckets (can hold up to loadFactor * 2^B items)
	//hash0     uint32 // hash seed
	//buckets   unsafe.Pointer
	//oldbuckets unsafe.Pointer
	//nevacuate  uint64
	//}

	// type bmap struct {
	//tophash [bucketCnt]uint8
	//data    [bucketCnt * 2]byte // 2 bytes per key/value pair
	//overflow unsafe.Pointer
	//}

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	println("Map value for 'one':", m["one"])
	println("Map value for 'two':", m["two"])

	sl := []int{1, 2, 3, 4, 5}
	//Slices:

	// A slice is a small header:
	// type SliceHeader struct {
	//	Data uintptr
	// Len  int
	// Cap  int
	// }

	println("Slice values:")
	for i, v := range sl {
		println(i, v)
	}
	println("Slice length:", len(sl))
}
