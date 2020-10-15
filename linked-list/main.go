// A singly linked list
package main

// TODO: convert to doubly linked list

type node struct {
	val interface{}
	next *node
}

type LinkedList struct {
	head *node
	len int
}

// Prepend adds a new node at the beginning of the linked list
func (ll *LinkedList) Prepend(v interface{}) {
	ll.head = &node{v, ll.head}

	ll.len++
}

// Append adds a new node at the end of the linked list
func (ll *LinkedList) Append(v interface{}) {
	n := &node{val: v}
	if ll.head == nil {
		ll.head = n
	} else {
		end := ll.head
		for {
			if end.next == nil {
				break
			}
			end = end.next
		}
		end.next = n
	}

	ll.len++
}

// Insert adds a new node at a specified point in the list
func (ll *LinkedList) Insert() {
	// TODO: Insert implementation
}

// Remove deletes a specified node in the list
func (ll *LinkedList) Remove() {
	// TODO: Remove implementation
}

// Clear empties the list of contents
func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.len = 0
}

// Size returns the length of the list
func (ll *LinkedList) Size() int {
	return ll.len
}

// IsEmpty returns whether or not the list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.len == 0
}

// Contents returns all contents in the list
func (ll *LinkedList) Contents() []interface{} {
	if ll.len == 0 {
		return nil
	}

	contents := make([]interface{}, ll.len)
	n := ll.head
	for n != nil {
		contents = append(contents, n.val)
		n = n.next
	}

	return contents
}
