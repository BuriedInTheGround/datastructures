package unionfind

import "fmt"

// UnionFind is a data structure that keeps track of elements which are split
// into one or more disjoint sets.
type UnionFind struct {
	data       []int
	size       []int
	components int
}

const (
	defaultUnionFindSize = 8
)

// New returns a new UnionFind instance.
//
// Complexity: O(n)
func New() UnionFind {
	return NewWithSize(defaultUnionFindSize)
}

// NewWithSize returns a new UnionFind instance with the specified `size` as
// number of elements.
//
// Complexity: O(n)
func NewWithSize(size int) UnionFind {
	if size <= 0 {
		panic("size must be a positive number")
	}

	data := make([]int, size)
	sz := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = i
		sz[i] = 1
	}

	return UnionFind{
		data:       data,
		size:       sz,
		components: size,
	}
}

// Size returns the number of elements that are into the UnionFind.
//
// Complexity: O(1)
func (uf *UnionFind) Size() int {
	return len(uf.data)
}

// Components returns the number of components (or groups) that the UnionFind
// has.
//
// Complexity: O(1)
func (uf *UnionFind) Components() int {
	return uf.components
}

// Find finds to which component/group the requested `element` belongs to and
// returns its root.
// This method also applies Path Compression.
//
// Complexity: O(α(n))
func (uf *UnionFind) Find(element int) int {
	if element > uf.Size() {
		panic("cannot exists such element inside this UnionFind")
	}

	root := element
	for root != uf.data[root] {
		root = uf.data[root]
	}

	uf.compressPath(element, root)

	return root
}

// Connected returns whether two elements belongs to the same component or
// not.
//
// Complexity: O(α(n))
func (uf *UnionFind) Connected(e1, e2 int) bool {
	return uf.Find(e1) == uf.Find(e2)
}

// ComponentSize returns the number of elements are in the same component as
// `element`.
//
// Complexity: O(α(n))
func (uf *UnionFind) ComponentSize(element int) int {
	return uf.size[uf.Find(element)]
}

// Unify performs the union operation that merges to components into one,
// making as new component root the root that previously was in the bigger
// component.
//
// Complexity: O(α(n))
func (uf *UnionFind) Unify(e1, e2 int) {
	root1 := uf.Find(e1)
	root2 := uf.Find(e2)

	if root1 == root2 {
		return
	}

	if uf.size[root1] < uf.size[root2] {
		uf.size[root2] += uf.size[root1]
		uf.data[root1] = root2
	} else {
		uf.size[root1] += uf.size[root2]
		uf.data[root2] = root1
	}

	uf.components--
}

func (uf *UnionFind) compressPath(from, root int) {
	for from != root {
		next := uf.data[from]
		uf.data[from] = root
		from = next
	}
}

func (uf UnionFind) String() string {
	res := "{ "
	for k, v := range uf.data {
		res += fmt.Sprintf("%d->%d ", k, v)
	}
	return res + "}"
}
