package main

import "fmt"

func main() {
	example(1000)
}

func example(i int) {
	arr := []int{}
	currCap := cap(arr)
	for j := 0; j < i; j++ {
		if cap(arr) != currCap {
			currCap = cap(arr)
			fmt.Println("New Capacity: ", currCap)
		}
		arr = append(arr, j)
	}
}

// no allocations needed since nothing could be used outside of the call stack
func main1() {
	_ = stackIt()
}

//go:noinline
func stackIt() int {
	y := 2
	return y * 2
}

// one allocation needed/op since the pointer is being passed back up the stack
func main2() {
	_ = stackIt2()
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

// no allocations needed since the pointer is only passed down the stack and isn't reused
func main3() {
	y := 2
	_ = stackIt3(&y) // pass y down the stack as a pointer
}

//go:noinline
func stackIt3(y *int) int {
	res := *y * 2
	return res
}
