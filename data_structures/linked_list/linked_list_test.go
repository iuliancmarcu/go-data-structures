package linked_list

import "testing"

func TestLinkedList_Insert(t *testing.T) {
	ll := LinkedList[int]{}

	// Test inserting values into an empty list
	ll.Insert(1)
	if ll.head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %v", ll.head.Value)
	}

	// Test inserting values into a non-empty list
	ll.Insert(2)
	ll.Insert(3)
	if ll.head.Next.Value != 2 {
		t.Errorf("Expected second node value to be 2, got %v", ll.head.Next.Value)
	}
	if ll.head.Next.Next.Value != 3 {
		t.Errorf("Expected third node value to be 3, got %v", ll.head.Next.Next.Value)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	// Test deleting a value from the middle of the list
	deleted := ll.Delete(2)
	if !deleted {
		t.Errorf("Expected value 2 to be deleted")
	}
	if ll.head.Next.Value != 3 {
		t.Errorf("Expected second node value to be 3 after deletion, got %v", ll.head.Next.Value)
	}

	// Test deleting the head value
	deleted = ll.Delete(1)
	if !deleted {
		t.Errorf("Expected value 1 to be deleted")
	}
	if ll.head.Value != 3 {
		t.Errorf("Expected head value to be 3 after deletion, got %v", ll.head.Value)
	}

	// Test deleting a value that doesn't exist in the list
	deleted = ll.Delete(5)
	if deleted {
		t.Errorf("Expected value 5 to not be deleted")
	}
}

func TestLinkedList_Has(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	// Test checking for a value that exists in the list
	has := ll.Has(2)
	if !has {
		t.Errorf("Expected value 2 to be found in the list")
	}

	// Test checking for a value that doesn't exist in the list
	has = ll.Has(5)
	if has {
		t.Errorf("Expected value 5 to not be found in the list")
	}
}
