package class_05_Linked_list

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (ll *LinkedList[T]) Display() error {
	if ll.head == nil {
		return errors.New("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v->", current.value)
		current = current.next
	}
	fmt.Println()
	return nil
}

func (ll *LinkedList[T]) Length() int {
	return ll.size
}
