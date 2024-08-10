package main

import (
	"cmp"
	"fmt"
)

func main() {
	list := &Node[int]{val: 1}
	list.Add(2)

	list.Add(3)
	list.Add(7)
	list.Insert(5, 2)
	list.Insert(4, 2)
	list.Insert(2, 20)
	current := list
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

	fmt.Println("i ", list.Index(7))
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

func (n *Node[T]) Add(val T) {
	curr := n
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node[T]{val: val}
}

func (n *Node[T]) Insert(val T, idx int) {
	if idx == 0 {
		nn := &Node[T]{val: val, next: n}
		*n = *nn
		return
	}

	curr := n
	for range idx - 1 {
		if curr.next == nil {
			fmt.Printf("index %d out of range\n", idx)
			return
		}
		curr = curr.next
	}
	nn := &Node[T]{val: val, next: curr.next}
	curr.next = nn

}

func (n *Node[T]) Index(val T) int {
	curr := n
	idx := 0
	for curr != nil {
		if curr.val == val {
			return idx
		}
		curr = curr.next
		idx++
	}

	return -1
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *TreeNode[T]
}
type TreeNode[T any] struct {
	val         T
	left, right *TreeNode[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *TreeNode[T]) Add(f OrderableFunc[T], v T) *TreeNode[T] {
	if n == nil {
		return &TreeNode[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}
func (n *TreeNode[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}
