package skip_list

import (
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	skipList := &SkipList[int, int]{}

	// Insert some values into the skip list
	skipList.Insert(1, 10)
	skipList.Insert(2, 20)
	skipList.Insert(3, 30)
	skipList.Insert(4, 0)

	// Verify that the values are inserted correctly
	if value := skipList.Find(1); value == nil || *value != 10 {
		t.Errorf("Expected value 10 for key 1, got %v", value)
	}
	if value := skipList.Find(2); value == nil || *value != 20 {
		t.Errorf("Expected value 20 for key 2, got %v", value)
	}
	if value := skipList.Find(3); value == nil || *value != 30 {
		t.Errorf("Expected value 30 for key 3, got %v", value)
	}
	if value := skipList.Find(4); value == nil || *value != 0 {
		t.Errorf("Expected value 0 for key 4, got %v", value)
	}
}

func TestSkipList_Delete(t *testing.T) {
	skipList := &SkipList[int, int]{}

	// Insert some values into the skip list
	skipList.Insert(1, 10)
	skipList.Insert(2, 20)
	skipList.Insert(3, 30)
	skipList.Insert(4, 0)

	// Delete a value from the skip list
	deleted := skipList.Delete(2)

	// Verify that the value is deleted
	if !deleted {
		t.Errorf("Expected value 20 to be deleted, but it was not")
	}

	// Verify that the value is no longer in the skip list
	if value := skipList.Find(2); value != nil {
		t.Errorf("Expected value 20 to be deleted, but it still exists in the skip list")
	}
}

func TestSkipList_Find(t *testing.T) {
	skipList := &SkipList[string, int]{}

	// Insert some values into the skip list
	skipList.Insert("1", 10)
	skipList.Insert("2", 20)
	skipList.Insert("3", 30)
	skipList.Insert("4", 0)

	// Find values in the skip list
	if value := skipList.Find("1"); value == nil || *value != 10 {
		t.Errorf("Expected value 10 for key 1, got %v", value)
	}
	if value := skipList.Find("2"); value == nil || *value != 20 {
		t.Errorf("Expected value 20 for key 2, got %v", value)
	}
	if value := skipList.Find("3"); value == nil || *value != 30 {
		t.Errorf("Expected value 30 for key 3, got %v", value)
	}
	if value := skipList.Find("4"); value == nil || *value != 0 {
		t.Errorf("Expected value 0 for key 4, got %v", value)
	}

	// Find a non-existent value in the skip list
	if value := skipList.Find("5"); value != nil {
		t.Errorf("Expected value to be nil for key 4, got %v", value)
	}
}
