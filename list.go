package main

/*
 * Simple generic double linked list. requires go 1.2.1 for cmp
 */
import (
	"cmp"
	"fmt"
)

type LinkedList[T cmp.Ordered] struct {
	head  *ListItem[T]
	tail  *ListItem[T]
	value T
}
type ListItem[T any] struct {
	next *ListItem[T]
	prev *ListItem[T]
	val  T
}

func (l *LinkedList[T]) Insert(t T) {
	if l.head == nil {
		l.head = &ListItem[T]{
			next: nil,
			prev: nil,
			val:  t,
		}
		l.tail = l.head
	} else {

		c := l.head

		for c.next != nil {
			c = c.next
		}
		c.next = &ListItem[T]{
			next: nil,
			prev: c,
			val:  t,
		}
		l.tail = c.next
	}

}

func (l *LinkedList[T]) Print() {
	p := l.head

	for p != nil {
		fmt.Printf("curr item: %v\n", p.val)
		p = p.next
	}
}
