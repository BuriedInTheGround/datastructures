package unionfind

import "testing"

func TestFind(t *testing.T) {
	uf := New()

	for i := 0; i < uf.Size(); i++ {
		if r := uf.Find(i); r != i {
			t.Errorf("wrong result: got %d want %d", r, i)
		}
	}
}

func TestConnected(t *testing.T) {
	uf := New()

	for i := 1; i < uf.Size(); i++ {
		if uf.Connected(i-1, i) {
			t.Errorf("elements `%d` and `%d` should not be connected", i-1, i)
		}
	}
}

func TestUnify(t *testing.T) {
	uf := New()

	uf.Unify(2, 3)
	if !uf.Connected(2, 3) {
		t.Errorf("union not working")
	}
	uf.Unify(1, 4)
	uf.Unify(2, 4)
	if !uf.Connected(1, 3) {
		t.Errorf("union not working")
	}

	if s := uf.ComponentSize(2); s != 4 {
		t.Errorf("wrong component size: got %d want %d", s, 4)
	}

	if c := uf.Components(); c != 5 {
		t.Errorf("wrong number of components: got %d want %d", c, 5)
	}

	if r := uf.Find(4); r != 2 {
		t.Errorf("wrong root: got %d want %d", r, 2)
	}
	t.Log(uf) // Here should be { 0->0 1->2 2->2 3->2 4->2 5->5 6->6 7->7 }.
}
