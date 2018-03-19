package main

import (
	"fmt"
)

type Tree struct {
	data int
	left *Tree
	right *Tree
}

func NewTree(data int) *Tree {
	return &Tree{data:data, left:nil, right:nil}
}

func (t *Tree) add(data int) *Tree {
	if t == nil {
		return NewTree(data)
	}
	if data <= t.data {
		t.left = t.left.add(data)
	} else {
		t.right = t.right.add(data)
	}
	return t
}

func (t *Tree) find(data int) *Tree {
	var result *Tree
	if t == nil {
		fmt.Println("Tree is nil")
		result = nil
	} else if data == t.data {
		fmt.Println("Found Node")
		result = t
	} else {
		if data < t.data {
			fmt.Println("Checking left branch")
			result = t.left.find(data)
		} else {
			fmt.Println("Checking right branch")
			result = t.right.find(data)
		}
	}
	return result
}

func (t *Tree) delete(data int) *Tree {
	return nil
}

func main() {
	fmt.Println("Tree")
	tree := NewTree(8)
	tree.add(3)
	tree.add(10)
	tree.add(1)
	tree.add(6)
	tree.add(4)
	tree.add(7)
	tree.add(14)
	tree.add(13)

	fmt.Printf("%+v\n", tree)
	fmt.Printf("%+v\n", tree.find(13))
	fmt.Printf("%+v\n", tree.find(6))
}
