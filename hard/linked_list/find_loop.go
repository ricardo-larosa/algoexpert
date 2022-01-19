package linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Tolerate acyclic and null inputs
func FindLoop(head *LinkedList) *LinkedList {
	var cycle bool // Use floyd cycle detection
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast { // found cycle
			cycle = true
			break // meeting point m
		}
	}
	// set slow to the beginning and iterate at same pace
	for slow = head; cycle; slow, fast = slow.Next, fast.Next {
		if slow == fast { // once they meet that's the start of the cycle
			return slow
		}
	}

	return nil
}
