package main

import (
	"cmp"
	"fmt"
)

func main() {
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	t2 := NewTree(OrderPeople)
	t2.Add(Person{Name: "Abrao", Age: 20})
	t2.Add(Person{Name: "Abrao", Age: 10})
	t2.Add(Person{Name: "Thiago", Age: 20})
	fmt.Println(t2.Contains(Person{Name: "Abrao", Age: 10}))
	fmt.Println(t2.Contains(Person{Name: "Abrao", Age: 11}))

	t3 := NewTree(Person.Order)
	t3.Add(Person{Name: "Abrao", Age: 20})
	t3.Add(Person{Name: "Abrao", Age: 10})
	t3.Add(Person{Name: "Thiago", Age: 20})
	fmt.Println(t3.Contains(Person{Name: "Abrao", Age: 10}))
	fmt.Println(t3.Contains(Person{Name: "Abrao", Age: 11}))

	list := &Node[int]{val: 1}
	list.Add(2)
	list.Add(3)
	list.Add(7)
	list.Insert(5, 2)
	fmt.Printf("idx: %d \n", list.Index(2))

	current := list
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
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
	curr := n
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	curr.val = val

}

func (n *Node[T]) Index(val T) int {
	curr := n
	idx := 0
	for curr.next != nil {
		if curr.val == val {
			return idx
		}
		curr = curr.next
		idx++
	}
	if curr.val == val {
		return idx
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
