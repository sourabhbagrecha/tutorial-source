package main

import "fmt"

const ArraySize = 7

//HashTable will
type HashTable struct {
	keyArray [ArraySize]*bucket
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

//Init will create a hash table and initialize the array with empty buckets
func Init() *HashTable {
	h := &HashTable{}
	for i := range h.keyArray {
		h.keyArray[i] = &bucket{}
	}
	return h
}

// Insert will take in a key and insert it in the hash table
func (h *HashTable) Insert(k string) {
	index := hash(k)
	h.keyArray[index].insert(k)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.keyArray[index].delete(k)
}

// Search will take in a key and return true if that key exists in the hash table
func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.keyArray[index].search(k)
}

func (b *bucket) insert(k string) {
	newNode := &bucketNode{data: k}
	if !b.search(k) {
		newNode.next = b.head
		b.head = newNode
		fmt.Println("node added", newNode.data)
	} else {
		fmt.Println(k, "is not added because it is already in the list")
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
	for currentNode != nil {
		if currentNode.data == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func main() {
	hashTable := Init()

	hashTable.Insert("RANDY")
	hashTable.Insert("ERIC")
	hashTable.Insert("KENNY")
	hashTable.Insert("KENNY")
	fmt.Println(hashTable.Search("KENNY"))
}
