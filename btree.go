package glinq

import g "github.com/yoyofx/glinq/generic"

const maxChildren = 64 // must be even and > 2

//BTree a B-tree data structures
type BTree[K, V any] struct {
	root   *BTreeNode[K, V]
	height int
	n      int

	less g.LessFn[K]
}

type BTreeNode[K, V any] struct {
	m        int
	children [maxChildren]entry[K, V]
}

type entry[K, V any] struct {
	key   K
	val   V
	valid bool
	next  *BTreeNode[K, V]
}

// NewBTree New returns an empty B-tree.
func NewBTree[K, V any](less g.LessFn[K]) *BTree[K, V] {
	return &BTree[K, V]{
		root: &BTreeNode[K, V]{},
		less: less,
	}
}

// Size returns the number of elements in the tree.
func (t *BTree[K, V]) Size() int {
	return t.n
}

// Get returns the value associated with 'key'.
func (t *BTree[K, V]) Get(key K) (V, bool) {
	return t.search(t.root, key, t.height)
}

func (t *BTree[K, V]) search(x *BTreeNode[K, V], key K, height int) (V, bool) {
	children := x.children

	if height == 0 {
		// leaf BTreeNode
		for j := 0; j < x.m; j++ {
			if g.Compare(key, children[j].key, t.less) == 0 {
				return children[j].val, children[j].valid
			}
		}
	} else {
		// internal BTreeNode
		for j := 0; j < x.m; j++ {
			if x.m == j+1 || g.Compare(key, children[j+1].key, t.less) < 0 {
				return t.search(children[j].next, key, height-1)
			}
		}
	}
	var v V
	return v, false
}

// Put associates 'key' with 'val'.
func (t *BTree[K, V]) Put(key K, val V) {
	u := t.insert(t.root, key, val, t.height, true)
	t.n++
	if u == nil {
		return
	}

	n := &BTreeNode[K, V]{
		m: 2,
	}
	n.children[0] = entry[K, V]{
		key:  t.root.children[0].key,
		next: t.root,
	}
	n.children[1] = entry[K, V]{
		key:  u.children[0].key,
		next: u,
	}
	t.root = n
	t.height++
}

// Remove removes the value associated with 'key'.
func (t *BTree[K, V]) Remove(key K) {
	_, ok := t.Get(key)
	if !ok {
		return
	}
	var v V
	// insert a tombstone to remove an existing value
	t.insert(t.root, key, v, t.height, false)
	t.n--
}

func (t *BTree[K, V]) insert(h *BTreeNode[K, V], key K, val V, height int, valid bool) *BTreeNode[K, V] {
	ent := entry[K, V]{
		key:   key,
		val:   val,
		valid: valid,
	}

	var j int
	if height == 0 {
		// leaf BTreeNode
		for j = 0; j < h.m; j++ {
			if g.Compare(key, h.children[j].key, t.less) == 0 {
				h.children[j].val = val
				h.children[j].valid = valid
				return nil
			} else if g.Compare(key, h.children[j].key, t.less) < 0 {
				break
			}
		}
	} else {
		// internal BTreeNode
		for j = 0; j < h.m; j++ {
			if (j+1 == h.m) || g.Compare(key, h.children[j+1].key, t.less) < 0 {
				u := t.insert(h.children[j].next, key, val, height-1, valid)
				if u == nil {
					return nil
				}
				j++
				ent.key = u.children[0].key
				ent.valid = false
				ent.next = u
				break
			}
		}
	}

	for i := h.m; i > j; i-- {
		h.children[i] = h.children[i-1]
	}
	h.children[j] = ent
	h.m++
	if h.m < maxChildren {
		return nil
	}
	return t.split(h)
}

func (t *BTree[K, V]) split(h *BTreeNode[K, V]) *BTreeNode[K, V] {
	n := &BTreeNode[K, V]{
		m: maxChildren / 2,
	}
	h.m = maxChildren / 2
	for j := 0; j < maxChildren/2; j++ {
		n.children[j] = h.children[maxChildren/2+j]
	}
	return n
}

// Each calls 'fn' on every BTreeNode in the tree in order.
func (t *BTree[K, V]) Each(fn func(key K, val V)) {
	t.each(t.root, t.height, fn)
}

func (t *BTree[K, V]) each(n *BTreeNode[K, V], height int, fn func(key K, val V)) {
	if height == 0 {
		for j := 0; j < n.m; j++ {
			if !n.children[j].valid {
				continue
			}
			fn(n.children[j].key, n.children[j].val)
		}
	} else {
		for j := 0; j < n.m; j++ {
			t.each(n.children[j].next, height-1, fn)
		}
	}
}
