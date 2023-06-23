package main

import (
	"fmt"
	"io"
)

type Node struct {
	data string
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Append(data string) *LinkedList {
	// Create a Node with the data
	node := &Node{
		data: data,
		Next: nil,
	}

	if l.Head == nil {
		// Edge case: LinkedList is empty, so this is the first node.
		l.Head = node
		l.Length = 1
	} else {
		// Traverse the list to the end
		n := l.Head
		for n.Next != nil {
			n = n.Next
		}

		// Append node
		n.Next = node
		l.Length++
	}

	return l
}

func (l *LinkedList) Print(w io.Writer) error {
	node := l.Head
	tmp := ""

	for node != nil {
		tmp += node.data
		if node.Next != nil {
			tmp += " -> "
		}

		node = node.Next
	}

	_, err := fmt.Fprintf(w, "LinkedList:%s\n", tmp)
	return err
}

// RemoveDuplicates accepts a LinkedList as a parameter and returns a
// LinkedList without duplicate values.
// The list passed in as a parameter should not be modified, a new list should be returned.
// Feel free to add any methods needed to the Node or LinkedList structs above.
func RemoveDuplicates(list *LinkedList) *LinkedList {
	if list == nil || list.Head == nil {
		return list
	}

	distinctValues := make(map[int]bool)
	distinctList := NewLinkedList()

	current := list.Head
	for current != nil {
		value := current.Value
		if !distinctValues[value] {
			distinctValues[value] = true
			distinctList.Add(value)
		}
		current = current.Next
	}

	return list
}
