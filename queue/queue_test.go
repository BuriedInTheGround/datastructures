package queue

import "testing"

func TestEnqueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.Size() != 5 {
		t.Errorf("queue should have a size of %d, but it does not", 5)
	}
}

func TestDequeue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for i := 1; !queue.IsEmpty(); i++ {
		if v := queue.Dequeue(); v != i {
			t.Errorf("wrong data: got %d want %d", v, i)
		}
	}
}

func TestPeek(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Dequeue()
	queue.Enqueue(3)
	queue.Dequeue()

	// Here should be ->[ 3 ]->

	if v := queue.Peek(); v != 3 {
		t.Errorf("wrong data: got %d want %d", v, 3)
	}
}

func TestContains(t *testing.T) {
	queue := New()

	if queue.Contains(42) {
		t.Errorf("the queue is empty, cannot contains any value")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	if !queue.Contains(2) {
		t.Errorf("queue must contains `%d`, but was not found", 2)
	}
	if queue.Contains(5) {
		t.Errorf("queue does not contains `%d`, but was found", 5)
	}
}

func TestRemoveFirstOccurrence(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.RemoveFirstOccurrence(3)

	for i := 1; !queue.IsEmpty(); i++ {
		if i == 3 {
			continue
		}
		if v := queue.Dequeue(); v != i {
			t.Errorf("wrong data: got %d want %d", v, i)
		}
	}
}
