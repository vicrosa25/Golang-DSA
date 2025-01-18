package class_05_Linked_list

import (
	"strconv"
	"strings"
)

const separator = "->"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func Serialize(node *Node) string {
	if node == nil {
		return ""
	}
	output := strconv.Itoa(node.Value) + separator
	for node.Next != nil {
		node = node.Next
		output += strconv.Itoa(node.Value) + separator
	}
	return strings.TrimSuffix(output, separator)
}

func Deserialize(stringRepresentation string) *Node {
	if stringRepresentation == "" {
		return nil
	}
	broken := strings.Split(stringRepresentation, separator)
	var curr, last *Node
	for i := len(broken) - 1; i >= 0; i-- {
		last = curr
		curr = &Node{
			Value: atoi(broken[i]),
			Next:  nil,
		}
		if last != nil {
			curr.Next = last
		}
	}
	return curr
}

func atoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		return -1
	}
	return i
}

func (ll *LinkedList) Length() int {
	return ll.size
}

func (ll *LinkedList) Prepend(value int) {
	node := &Node{Value: value}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		node.Next = ll.head
		ll.head = node
	}
	ll.size++
}
