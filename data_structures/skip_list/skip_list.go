package skip_list

import (
	"cmp"
	"fmt"
	"math"
	"math/rand"
)

const P = 0

func should_grow() bool {
	return rand.Float32() < P
}

func random_level() int {
	level := 0
	for should_grow() {
		level++
	}
	return level
}

type NodeInfo[K cmp.Ordered, T any] struct {
	Node     *Node[K, T]
	Distance uint
}

type Node[K cmp.Ordered, T any] struct {
	Value T
	Key   K
	Next  []*NodeInfo[K, T]
}

func (n *Node[K, T]) String() string {
	return fmt.Sprintf("[Node] level: %v value %v", n.Level(), n.Value)
}

func (n *Node[K, T]) Level() int {
	return len(n.Next) - 1
}

type SkipList[K cmp.Ordered, T any] struct {
	Head *Node[K, T]
}

func (l *SkipList[K, T]) Debug() {
	for level := l.Head.Level(); level >= 0; level-- {
		current_node := l.Head
		for current_node != nil {
			if current_node.Level() >= level {
				if current_node.Next[level] != nil {
					fmt.Printf("%v(%v)\t", current_node.Value, current_node.Next[level].Distance)
				} else {
					fmt.Printf("%v\t", current_node.Value)
				}
			} else {
				fmt.Print("\t")
			}

			if current_node.Next[0] != nil {
				current_node = current_node.Next[0].Node
			} else {
				current_node = nil
			}
		}

		fmt.Println()
	}
	fmt.Println()
}

func (l *SkipList[K, T]) Insert(key K, value T) {
	new_node_level := random_level()

	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		l.Head = &Node[K, T]{Next: make([]*NodeInfo[K, T], new_node_level+1)}
	}

	max_level := int(math.Max(float64(l.Head.Level()), float64(new_node_level)))

	current_node := l.Head
	nodes_to_update := make([]*Node[K, T], max_level+1)
	distance_to_node := make([]uint, max_level+1)

	// For each level from list Head top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		distance := uint(0)
		for current_node.Next[level] != nil && current_node.Next[level].Node.Key <= key {
			if current_node != l.Head {
				distance++
			}
			current_node = current_node.Next[level].Node
		}

		// Store the stopping node to be updated
		nodes_to_update[level] = current_node
		distance_to_node[level] = distance
	}
	fmt.Printf("%v\n", distance_to_node)

	// Create a new node
	node_to_insert := &Node[K, T]{
		Key: key, Value: value,
		Next: make([]*NodeInfo[K, T], new_node_level+1),
	}

	// Mark for update higher levels of Head if necessary
	if l.Head.Level() < new_node_level {
		for level := new_node_level; level > l.Head.Level(); level-- {
			nodes_to_update[level] = l.Head
		}

		// And expand the Head Next slice
		l.Head.Next = append(make([]*NodeInfo[K, T], 0, max_level+1), l.Head.Next...)
	}

	// Insert the node by updating the links
	level_distance := uint(0)
	for level := 0; level <= new_node_level; level++ {
		node_to_update := nodes_to_update[level]
		level_distance += distance_to_node[level]

		if node_to_update.Level() >= level {
			if node_to_update.Next[level] != nil {
				node_to_insert.Next[level] = &NodeInfo[K, T]{
					Node:     node_to_update.Next[level].Node,
					Distance: node_to_update.Next[level].Distance - level_distance,
				}
			}
			node_to_update.Next[level] = &NodeInfo[K, T]{
				Node:     node_to_insert,
				Distance: level_distance,
			}
		}
	}
}

func (l *SkipList[K, T]) Delete(key K) bool {
	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		return false
	}

	current_node := l.Head
	nodes_to_update := make([]*Node[K, T], l.Head.Level()+1)
	distance_to_node := make([]uint, l.Head.Level()+1)

	// For each level from top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		distance := uint(0)
		for current_node.Next[level] != nil && current_node.Next[level].Node.Key <= key {
			current_node = current_node.Next[level].Node
			distance++
		}

		// And save node to update later
		nodes_to_update[level] = current_node
		distance_to_node[level] = distance
	}

	// Next node has to be be the node containing the value
	node_to_delete := current_node.Next[0].Node
	if node_to_delete.Key != key {
		return false
	}

	// Delete the node by updating the saved nodes
	level_distance := uint(0)
	for level := 0; level <= l.Head.Level(); level++ {
		node_to_update := nodes_to_update[level]
		level_distance += distance_to_node[level]

		if node_to_delete.Level() >= level && node_to_update.Next[level] != nil {
			node_to_update.Next[level] = &NodeInfo[K, T]{
				Node:     node_to_delete.Next[level].Node,
				Distance: node_to_update.Next[level].Distance + level_distance,
			}
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

func (l *SkipList[K, T]) Find(key K) *T {
	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		return nil
	}

	current_node := l.Head

	// For each level from top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		for current_node.Next[level] != nil && current_node.Next[level].Node.Key <= key {
			current_node = current_node.Next[level].Node
		}

		if current_node.Key == key {
			return &current_node.Value
		}
	}

	return nil
}

func (l *SkipList[K, T]) At(index uint) *T {
	// FIXME: Safeguard operation for uninitialized list
	if l.Head == nil {
		return nil
	}

	current_node := l.Head
	distance := uint(0)

	// For each level from top-most
	for level := l.Head.Level(); level >= 0; level-- {
		// Move right as much as possible
		for current_node.Next[level] != nil && distance+current_node.Next[level].Distance <= index {
			distance += current_node.Next[level].Distance
			current_node = current_node.Next[level].Node
		}

		if distance == index {
			return &current_node.Value
		}
	}

	return nil
}
