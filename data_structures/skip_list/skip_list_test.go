package skip_list

import (
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	skipList := &SkipList[int]{}

	// Insert values into the skip list
	skipList.Insert(5)
	skipList.Insert(10)
	skipList.Insert(3)
	skipList.Insert(7)
	skipList.Insert(2)
	skipList.Insert(19)
	skipList.Insert(0)

	// Check if the values are present in the skip list
	if !skipList.Has(5) {
		t.Errorf("Expected value 5 to be present in the skip list")
	}
	if !skipList.Has(10) {
		t.Errorf("Expected value 10 to be present in the skip list")
	}
	if !skipList.Has(3) {
		t.Errorf("Expected value 3 to be present in the skip list")
	}
	if !skipList.Has(7) {
		t.Errorf("Expected value 7 to be present in the skip list")
	}
}

func TestSkipList_Delete(t *testing.T) {
	skipList := &SkipList[int]{}

	// Insert values into the skip list
	skipList.Insert(5)
	skipList.Insert(10)
	skipList.Insert(3)
	skipList.Insert(7)
	skipList.Insert(2)
	skipList.Insert(19)
	skipList.Insert(0)

	// Delete values from the skip list
	skipList.Delete(5)
	skipList.Delete(10)

	// Check if the deleted values are no longer present in the skip list
	if skipList.Has(5) {
		t.Errorf("Expected value 5 to be deleted from the skip list")
	}
	if skipList.Has(10) {
		t.Errorf("Expected value 10 to be deleted from the skip list")
	}

	// Check if the remaining values are still present in the skip list
	if !skipList.Has(3) {
		t.Errorf("Expected value 3 to be present in the skip list")
	}
	if !skipList.Has(7) {
		t.Errorf("Expected value 7 to be present in the skip list")
	}
}

func TestSkipList_Has(t *testing.T) {
	skipList := &SkipList[int]{}

	// Insert values into the skip list
	skipList.Insert(5)
	skipList.Insert(10)
	skipList.Insert(3)
	skipList.Insert(7)

	// Check if the values are present in the skip list
	if !skipList.Has(5) {
		t.Errorf("Expected value 5 to be present in the skip list")
	}
	if !skipList.Has(10) {
		t.Errorf("Expected value 10 to be present in the skip list")
	}
	if !skipList.Has(3) {
		t.Errorf("Expected value 3 to be present in the skip list")
	}
	if !skipList.Has(7) {
		t.Errorf("Expected value 7 to be present in the skip list")
	}

	// Check if a non-existent value is not present in the skip list
	if skipList.Has(15) {
		t.Errorf("Expected value 15 to not be present in the skip list")
	}
}
