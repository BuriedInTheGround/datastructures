package doublylinkedlist

import "testing"

func TestInsertFromHead(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(3)
	linkedList.InsertFromHead(4)
	linkedList.InsertFromHead(5)

	n := linkedList.Head()
	for i := 5; n != nil; i-- {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	for i := 1; n != nil; i++ {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Prev()
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
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	for i := 5; n != nil; i-- {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Prev()
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
	wantForward := []int{5, 3, 1, 2, 4}
	for i := 0; n != nil; i++ {
		if n.Content() != wantForward[i] {
			t.Errorf("wrong data: got %d want %d", n.Content(), wantForward[i])
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	wantBackward := []int{4, 2, 1, 3, 5}
	for i := 0; n != nil; i++ {
		if n.Content() != wantBackward[i] {
			t.Errorf("wrong data: got %d want %d", n.Content(), wantBackward[i])
		}
		n = n.Prev()
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
	linkedList.RemoveFromHead()
	linkedList.InsertFromTail(4)
	linkedList.RemoveFromHead()
	linkedList.InsertFromHead(2)
	linkedList.InsertFromTail(5)
	linkedList.InsertFromHead(1)

	// Should now be [ 1, 2, 3, 4, 5 ].

	n := linkedList.Head()
	for i := 1; n != nil; i++ {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	for i := 5; n != nil; i-- {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Prev()
	}
}

func TestRemoveFromTail(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(3)
	linkedList.RemoveFromTail()
	linkedList.InsertFromHead(4)
	linkedList.RemoveFromTail()
	linkedList.InsertFromTail(2)
	linkedList.InsertFromHead(5)
	linkedList.InsertFromTail(1)

	// Should now be [ 5, 4, 3, 2, 1 ].

	n := linkedList.Head()
	for i := 5; n != nil; i-- {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	for i := 1; n != nil; i++ {
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Prev()
	}
}

func TestRemoveFirstOccurrence(t *testing.T) {
	linkedList := New()

	linkedList.InsertFromHead(1)
	linkedList.InsertFromHead(2)
	linkedList.InsertFromHead(3)
	linkedList.InsertFromHead(4)
	linkedList.InsertFromHead(5)

	if err := linkedList.RemoveFirstOccurrence(42); err == nil {
		t.Errorf("value %d is not in the list but was found", 42)
	}

	if err := linkedList.RemoveFirstOccurrence(3); err != nil {
		t.Errorf("value %d to remove not found in the list", 3)
	}

	n := linkedList.Head()
	for i := 5; n != nil; i-- {
		if i == 3 {
			continue
		}
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Next()
	}

	n = linkedList.Tail()
	for i := 1; n != nil; i++ {
		if i == 3 {
			continue
		}
		if n.Content() != i {
			t.Errorf("wrong data: got %d want %d", n.Content(), i)
		}
		n = n.Prev()
	}

	if linkedList.Size() != 4 {
		t.Errorf("wrong sze, the list should contains %d elements", 4)
	}
}
