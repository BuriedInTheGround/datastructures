package binarysearchtree

import (
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	var err error
	bst := New()

	for _, v := range []int{4, 3, 5, 6, 7, 1, 2} {
		err = bst.Insert(v)
		if err != nil {
			t.Errorf("insert returned error, but should not")
		}
	}

	for _, v := range []int{4, 3, 5, 6, 7, 1, 2} {
		err = bst.Insert(v)
		if err == nil {
			t.Errorf("insert should have returned an error")
		}
	}

	if s := bst.Size(); s != 7 {
		t.Errorf("wrong size: got %d want %d", s, 7)
	}

	inOrder := bst.TraverseInOrder()
	for i, v := range inOrder {
		if v != i+1 {
			t.Errorf("wrong traversal: got %d want %d", v, i+1)
		}
	}
}

func TestContains(t *testing.T) {
	bst := New()

	for _, v := range []int{4, 3, 5, 6, 7, 1, 2} {
		bst.Insert(v)
	}

	for i := 1; i <= bst.Size(); i++ {
		if !bst.Contains(i) {
			t.Errorf("the tree should contains %d, but tells it does not", i)
		}
	}
}

func TestRemove(t *testing.T) {
	bst := New()

	for _, v := range []int{4, 3, 5, 6, 8, 7, 1, 2} {
		bst.Insert(v)
	}

	toRemove := []int{5, 3, 7}
	for _, r := range toRemove {
		bst.Remove(r)
	}

	if h := bst.Height(); h != 3 {
		t.Errorf("wrong height: got %d want %d", h, 3)
	}

	inOrder := bst.TraverseInOrder()
	sort.Ints(toRemove)
	for i, v := range inOrder {
		for _, r := range toRemove {
			if i+1 >= r {
				i++
			}
		}
		if v != i+1 {
			t.Errorf("wrong traversal: got %d want %d", v, i+1)
		}
	}
}
