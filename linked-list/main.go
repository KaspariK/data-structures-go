// A singly linked list
package main

type node struct {
	val interface{}
	next *node
}

type linkedList struct {
	head *node
	len int
}

// Add new node at the beginning of the linked list
func (ll *linkedList) prepend(n *node) {
	old := ll.head
	ll.head = n
	ll.head.next = old

	ll.len++
}

// Add a new node at the end of the linked list
func (ll *linkedList) append(n *node) {
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

func (ll *linkedList) insert(v interface{}, n *node) {
	// TODO
}

func (ll *linkedList) remove(n *node) {
	// TODO
}

func (ll *linkedList) size() int {
	return ll.len
}

func (ll *linkedList) isEmpty() bool {
	return ll.len == 0
}


