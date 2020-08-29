package main

import "fmt"

const ArraySize = 7

//HashTable will
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket will
type bucket struct {
	head *bucketNode
}

// bucketNode will
type bucketNode struct {
	data string
	next *bucketNode
}

//hash will take in a string and return a hash code
func hash(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v)
	}
	return sum % ArraySize
}

// Insert will take in a key and insert it in the hash table
func (h *HashTable) Insert(k string) {
	// go to that array index
	index := hash(k)
	// insert the key in the bucket
	if h.array[index] == nil {
		h.array[index] = &bucket{}
	}
	h.array[index].insert(k)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

// Search will take in a key and return true if that key exists in the hash table
func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.array[index].search(k)
}

func (b *bucket) insert(k string) {
	// create bucketNode to insert
	newNode := &bucketNode{data: k}
	// insert that bucketNode
	if b.head == nil {
		b.head = newNode
		fmt.Println("node added", newNode.data)
	} else {
		newNode.next = b.head
		b.head = newNode
		fmt.Println("node added", newNode.data)
	}
}

func (b *bucket) delete(k string) {
	if b.head.data == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next.data != k {
		if previousNode.next.next == nil {
			return
		}
		previousNode = previousNode.next
	}
	previousNode.next = previousNode.next.next
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode.next != nil {
		if currentNode.data == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func main() {
	hashTable := &HashTable{}
	hashTable.Insert("RANDY")
	hashTable.Insert("ERIC")
	fmt.Println(hashTable.Search("rt"))
}
