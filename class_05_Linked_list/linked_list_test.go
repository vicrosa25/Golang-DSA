package class_05_Linked_list

import "testing"

/*
TestSerializeAndDeserializeLinkedList tests solution(s) with the following signature and problem description:

	type Node struct {
		Value int
		Next *Node
	}

	func Serialize(node *Node) string
	func Deserialize(stringRepresentation string) *Node

Write a function that turns a linked list into a string representation (Serialize), and then a function
that turns that string representation to an actual linked list (Deserialize).
*/

func TestSerializeAndDeserializeLinkedList(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1->2",
		"1->2->3",
		"1->2->3->4->5",
	}
	for i, test := range tests {
		got := Serialize(Deserialize(test))
		if got != test {
			t.Fatalf("Failed test case %d: want %#v, got %#v", i, test, got)
		}
	}
}
