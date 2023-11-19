package skip_list

import (
	"cmp"
	"fmt"
	"math"
	"math/rand"
)

const P = 0.75

type Node[T cmp.Ordered] struct {
	Value T
	Next  []*Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("[Node] level: %v value %v", n.Level(), n.Value)
}

func (n *Node[T]) Level() int {
	return len(n.Next) - 1
}

func flip_coin() bool {
	return rand.Float32() < P
}

func random_level() int {
	level := 0
	for flip_coin() {
		level++
	}
	return level
}

type SkipList[T cmp.Ordered] struct {
	Head *Node[T]
}

func (l *SkipList[T]) Debug() {
	for level := l.Head.Level(); level >= 0; level-- {
		current_node := l.Head
		for current_node != nil {
			if current_node.Level() >= level {
				fmt.Print(current_node.Value)
				fmt.Print("\t")
			} else {
				fmt.Print("\t")
			}

			current_node = current_node.Next[0]
		}

		fmt.Println()
	}
	fmt.Println()
}

func (l *SkipList[T]) Insert(value T) {
	new_node_level := random_level()

	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		l.Head = &Node[T]{Next: make([]*Node[T], new_node_level+1)}
	}

	max_level := int(math.Max(float64(l.Head.Level()), float64(new_node_level)))

	current_node := l.Head
	nodes_to_update := make([]*Node[T], max_level+1) // <-- +1 because we need cap() to include index 0

	// For each level from list Head top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		for current_node.Next[level] != nil && current_node.Next[level].Value <= value {
			current_node = current_node.Next[level]
		}

		// Store the stopping node to be updated
		nodes_to_update[level] = current_node
	}

	// Create a new node
	node_to_insert := &Node[T]{Value: value, Next: make([]*Node[T], new_node_level+1)} // <-- +1 because we need cap() to include index 0

	// Mark for update higher levels of Head if necessary
	if l.Head.Level() < new_node_level {
		for lvl := new_node_level; lvl > l.Head.Level(); lvl-- {
			nodes_to_update[lvl] = l.Head
		}

		// And expand the Head Next slice
		l.Head.Next = append(make([]*Node[T], 0, max_level+1), l.Head.Next...)
	}

	// Insert the node by updating the links
	for level := new_node_level; level >= 0; level-- {
		node_to_update := nodes_to_update[level]
		if node_to_update.Level() >= level {
			node_to_insert.Next[level] = node_to_update.Next[level]
			node_to_update.Next[level] = node_to_insert
		}
	}
}

func (l *SkipList[T]) Delete(value T) bool {
	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		return false
	}

	current_node := l.Head
	nodes_to_update := make([]*Node[T], l.Head.Level()+1)

	// For each level from top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		for current_node.Next[level] != nil && current_node.Next[level].Value < value {
			current_node = current_node.Next[level]
		}

		// And save node to update later
		nodes_to_update[level] = current_node
	}

	// Next node has to be be the node containing the value
	node_to_delete := current_node.Next[0]
	if node_to_delete.Value != value {
		return false
	}

	// Delete the node by updating the saved nodes
	for level := l.Head.Level(); level >= 0; level-- {
		node_to_update := nodes_to_update[level]
		if node_to_delete.Level() >= level {
			node_to_update.Next[level] = node_to_delete.Next[level]
		}
	}

	// Find the max level with link from Head
	new_max_level := 0
	for level, node := range l.Head.Next {
		if node != nil {
			new_max_level = level
		}
	}

	// Decrease the size of Head Next list if necessary
	if new_max_level < l.Head.Level() {
		// And reduce the Head Next slice to new max_level
		l.Head.Next = l.Head.Next[: new_max_level+1 : new_max_level+1]
	}

	return true
}

func (l *SkipList[T]) Has(value T) bool {
	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		return false
	}

	current_node := l.Head

	// For each level from top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		for current_node.Next[level] != nil && current_node.Next[level].Value <= value {
			current_node = current_node.Next[level]
		}

		if current_node.Value == value {
			return true
		}
	}

	return false
}
