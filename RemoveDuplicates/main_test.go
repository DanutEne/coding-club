package main

import "os"

func ExampleRemoveDuplicates() {
	l1 := &LinkedList{}
	l1 = l1.Append("hello").Append("world").Append("world").Append("hello")

	l2 := &LinkedList{}
	l2.Append("example").Append("linkedlists").Append("are").Append("fun")

	l3 := &LinkedList{}

	uniqL1 := RemoveDuplicates(l1) // uniq should contain the following nodes: Node{"hello"} -> Node{"world"}
	_ = l1.Print(os.Stdout)
	_ = uniqL1.Print(os.Stdout)

	uniqL2 := RemoveDuplicates(l2) // no duplicates, should output the same
	_ = l2.Print(os.Stdout)
	_ = uniqL2.Print(os.Stdout)

	uniqL3 := RemoveDuplicates(l3) // Empty Linked List shouldn't crash :)
	_ = l3.Print(os.Stdout)
	_ = uniqL3.Print(os.Stdout)

	// Output:
	// LinkedList:hello -> world -> world -> hello
	// LinkedList:hello -> world
	// LinkedList:example -> linkedlists -> are -> fun
	// LinkedList:example -> linkedlists -> are -> fun
	// LinkedList:
	// LinkedList:
}
