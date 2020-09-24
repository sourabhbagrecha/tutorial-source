package main

import (
	"fmt"
	"log"
)

// MaxHeap is a slice that holds the array
type MaxHeap struct {
	array []int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(i int) {

	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

//Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	l := len(h.array) - 1
	if l == 0 {
		log.Printf("cannot extract because array length is 0")
		return -1
	}
	extracted := h.array[0]
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyDown(i int) {

	lastIndex := len(h.array) - 1
	l, r := left(i), right(i)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when i only has one child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if current is smaller
		if h.array[i] < h.array[childToCompare] {
			h.swap(i, childToCompare)
			//update loop
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			// when the i is finally larger than the larger child
			return
		}

	}

}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}
