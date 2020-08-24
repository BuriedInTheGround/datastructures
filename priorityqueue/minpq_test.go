package priorityqueue

import "testing"

func TestAdd(t *testing.T) {
	pq := New()

	pq.Add(2)
	pq.Add(5)
	pq.Add(4)
	pq.Add(1)
	pq.Add(3)

	if pq.Size() != 5 {
		t.Errorf("the PQ should have a size of %d, but it does not", 5)
	}
}

func TestRemoveMin(t *testing.T) {
	pq := New()

	pq.Add(2)
	pq.Add(5)
	pq.Add(4)
	pq.Add(1)
	pq.Add(3)

	for i := 1; !pq.IsEmpty(); i++ {
		if v := pq.RemoveMin(); v != i {
			t.Errorf("wrong data: got %d want %d", v, i)
		}
	}
}

func TestPeek(t *testing.T) {
	pq := New()

	pq.Add(2)
	pq.Add(5)
	pq.Add(4)
	pq.Add(1)
	pq.Add(3)

	if v := pq.Peek(); v != 1 {
		t.Errorf("wrong data: got %d want %d", v, 1)
	}
}

func TestContains(t *testing.T) {
	pq := New()

	if pq.Contains(42) {
		t.Errorf("the PQ is empty, cannot contains any value")
	}

	pq.Add(2)
	pq.Add(4)
	pq.Add(1)
	pq.Add(3)

	if !pq.Contains(2) {
		t.Errorf("the PQ must contains `%d`, but was not found", 2)
	}
	if pq.Contains(5) {
		t.Errorf("the PQ does not contains `%d`, but was found", 5)
	}
}

func TestRemoveFirstOccurrence(t *testing.T) {
	pq := New()

	if err := pq.RemoveFirstOccurrence(42); err == nil {
		t.Errorf("the PQ is empty, cannot contains any value")
	}

	pq.Add(2)
	pq.Add(5)
	pq.Add(4)
	pq.Add(1)
	pq.Add(3)
	pq.Add(7)
	pq.Add(11)
	pq.Add(8)
	pq.Add(6)
	pq.Add(10)
	pq.Add(9)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			err := pq.RemoveFirstOccurrence(i)
			if err != nil {
				t.Errorf("error while removing valid value %d", i)
			}
		}
	}
	for i := 1; !pq.IsEmpty(); i += 2 {
		if v := pq.RemoveMin(); v != i {
			t.Errorf("wrong data: got %d want %d", v, i)
		}
	}
}
