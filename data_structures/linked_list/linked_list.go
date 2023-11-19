package linked_list

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Insert(value T) {
	new_node := Node[T]{Value: value}

	if l.head == nil {
		l.head = &new_node
		return
	}

	last_node := l.head
	for last_node.Next != nil {
		last_node = last_node.Next
	}

	last_node.Next = &new_node
}

func (l *LinkedList[T]) Delete(value T) bool {
	if l.head.Value == value {
		l.head = l.head.Next
		return true
	}

	current_node := l.head
	for current_node.Next != nil {
		if current_node.Next.Value == value {
			current_node.Next = current_node.Next.Next
			return true
		}

		current_node = current_node.Next
	}

	return false
}

func (l *LinkedList[T]) Has(value T) bool {
	current_node := l.head
	for current_node != nil {
		if current_node.Value == value {
			return true
		}
		current_node = current_node.Next
	}

	return false
}
