package main

import (
	"fmt"

	"github.com/esferadigital/ringbuffer"
)

func main() {
	// Create a ring buffer of capacity 3
	r := ringbuffer.New[int](3)

	fmt.Println("---- filling buffer with values ----")
	r.Write(10)
	r.Write(20)
	r.Write(30)

	fmt.Printf("len = %d\ncap = %d\nsnapshot = %v\n", r.Len(), r.Cap(), r.Snapshot())

	fmt.Println("\n---- write '40' when full ----")
	r.Write(40)
	fmt.Printf("len = %d\ncap = %d\nsnapshot = %v\n", r.Len(), r.Cap(), r.Snapshot())

	fmt.Println("\n---- peek the oldest element ----")
	if v, ok := r.Peek(); ok {
		fmt.Printf("peek = %d\n", v)
	}

	fmt.Println("\n---- read all elements ----")
	for {
		v, ok := r.Read()
		if !ok {
			break
		}
		fmt.Printf("read = %d\n", v)
	}
	fmt.Printf("len = %d\n", r.Len())

	fmt.Println("\n---- write '99' and '100' after emptying ----")
	r.Write(99)
	r.Write(100)
	fmt.Printf("snapshot now = %v\n", r.Snapshot())

	fmt.Println("\n---- reset the buffer ----")
	r.Reset()
	fmt.Printf("len = %d\nsnapshot = %v\n", r.Len(), r.Snapshot())
}
