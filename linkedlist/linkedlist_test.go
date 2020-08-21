package linkedlist

import (
	"testing"
)

func TestInsertFromHead(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(3)
	linkedList.InsertFromHead(4)
	linkedList.InsertFromHead(5)

	n := linkedList.Head()
	for i := 5; n != nil; i-- {
		if n.content != i {
			t.Errorf("wrong data: got %d want %d", n.content, i)
		}
		n = n.next
	}
}

func TestInsertFromTail(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromTail(1)
	linkedList.InsertFromTail(2)
	linkedList.InsertFromTail(3)
	linkedList.InsertFromTail(4)
	linkedList.InsertFromTail(5)

	n := linkedList.Head()
	for i := 1; n != nil; i++ {
		if n.content != i {
			t.Errorf("wrong data: got %d want %d", n.content, i)
		}
		n = n.next
	}
}

func TestInsertMixed(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromTail(2)
	linkedList.InsertFromHead(3)
	linkedList.InsertFromTail(4)
	linkedList.InsertFromHead(5)

	n := linkedList.Head()
	want := []int{5, 3, 1, 2, 4}
	for i := 0; n != nil; i++ {
		if n.content != want[i] {
			t.Errorf("wrong data: got %d want %d", n.content, want[i])
		}
		n = n.next
	}
}

func TestContains(t *testing.T) {
	linkedList := New()

	if linkedList.Contains(42) {
		t.Errorf("the list is empty, but tells it is not")
	}

	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(4)
	linkedList.InsertFromTail(6)
	linkedList.InsertFromTail(8)
	linkedList.InsertFromHead(10)

	for i := 1; i < 12; i++ {
		// The list should contains only even elements.
		if linkedList.Contains(i) != (i%2 == 0) {
			t.Errorf("the list should contains %d, but tells it does not", i)
		}
	}
}

func TestRemoveFromHead(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromTail(1)
	linkedList.InsertFromTail(2)
	linkedList.InsertFromTail(3)
	linkedList.InsertFromTail(4)
	linkedList.InsertFromTail(5)

	for i := 1; i <= 5; i++ {
		if removed := linkedList.RemoveFromHead(); removed != i {
			t.Errorf("wrong data removed: got %d want %d", removed, i)
		}
		t.Log(linkedList)
	}

	if !linkedList.IsEmpty() {
		t.Errorf("wrong size, the list should be empty")
	}
}

func TestRemoveFromTail(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(3)
	linkedList.InsertFromHead(4)
	linkedList.InsertFromHead(5)

	for i := 1; i <= 5; i++ {
		if removed := linkedList.RemoveFromTail(); removed != i {
			t.Errorf("wrong data removed: got %d want %d", removed, i)
		}
		t.Log(linkedList)
	}

	if !linkedList.IsEmpty() {
		t.Errorf("wrong size, the list should be empty")
	}
}
