package algo

import (
	"algo/util"
	"errors"
)

type Heap interface {
	Add(int)
	Min() int
	Pop() int
	//Change(int, int)
	Size() int
}

type BinaryHeap struct {
	values []int
}

func BuildBinaryHeap(values []int) BinaryHeap {
	result := util.ArrayCopy(values)
	for i := len(result)/2 + 1; i >= 0; i-- {
		siftDown(result, i)
	}

	return BinaryHeap{
		values: result,
	}
}

func Merge(h1, h2 BinaryHeap) BinaryHeap {
	len1 := len(h1.values)
	len2 := len(h2.values)
	newValues := make([]int, len1+len2)

	for i := 0; i < len1; i++ {
		newValues[i] = h1.values[i]
	}

	for i := 0; i < len2; i++ {
		newValues[i+len1] = h2.values[i]
	}

	return BuildBinaryHeap(newValues)
}

func (h BinaryHeap) Add(val int) {
	h.values = append(h.values, val)
	siftUp(h.values, len(h.values)-1)
}

func (h BinaryHeap) Min() int {
	return h.values[0]
}

func (h BinaryHeap) Pop() int {
	result := h.values[0]

	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[0 : len(h.values)-1]
	siftDown(h.values, 0)

	return result
}

func (h BinaryHeap) Change(idx, new int) {
	old := h.values[idx]

	if new == old {
		return
	}

	h.values[idx] = new

	if new > old {
		siftUp(h.values, idx)
	} else {
		siftDown(h.values, idx)
	}
}

func (h BinaryHeap) Size() int {
	return len(h.values)
}

func siftDown(array []int, pos int) {
	len := len(array)

	left := (2 * pos) + 1
	right := (2 * pos) + 2

	for left < len {
		idx := pos
		val := array[pos]

		if left < len {
			if val > left {
				idx = left
				val = array[idx]
			}
		}

		if right < len {
			if val > right {
				idx = right
				val = array[idx]
			}
		}

		if idx != pos {
			t := array[pos]
			array[pos] = array[idx]
			array[idx] = t
		} else {
			break
		}

		pos = idx
		left = (2 * pos) + 1
		right = (2 * pos) + 2
	}
}

func siftUp(array []int, pos int) {
	p := parent(pos)

	for p >= 0 {
		if array[p] > array[pos] {
			t := array[pos]
			array[pos] = array[p]
			array[p] = t
		} else {
			break
		}
		pos = p
		p = parent(pos)
	}
}

func parent(idx int) int {
	if (idx % 2) == 0 {
		idx -= 2
	} else {
		idx--
	}

	return idx / 2
}

type BinominalHeap struct {
	root *BinominalNode
}

type BinominalNode struct {
	Key     int
	Parent  *BinominalNode
	Sibling *BinominalNode
	Child   *BinominalNode
	Degree  int
}

func BuildBinominalHeap(arr []int) BinominalHeap {
	if len(arr) == 0 {
		return BinominalHeap{}
	}

	heap := BinominalHeap{
		root: &BinominalNode{Key: arr[0], Child: nil, Parent: nil, Sibling: nil, Degree: 0},
	}

	for i := 1; i < len(arr); i++ {
		heap.Add(arr[i])
	}

	return heap
}

func (h BinominalHeap) Add(val int) {
	heap := BinominalHeap{
		root: &BinominalNode{
			Key:     val,
			Parent:  nil,
			Child:   nil,
			Sibling: nil,
			Degree:  0,
		},
	}

	h.root = merge(h.root, heap.root)
}

func (h BinominalHeap) Min() int {
	if h.root == nil {
		panic(errors.New("heap is empty"))
	}
	min := h.root.Key
	cur := h.root.Sibling
	for cur != nil {
		if min > cur.Key {
			min = cur.Key
		}
		cur = cur.Sibling
	}

	return min
}

func (h BinominalHeap) Pop() int {
	if h.root == nil {
		panic(errors.New("heap is empty"))
	}

	var prev *BinominalNode = nil
	var minPrev *BinominalNode = nil
	min := h.root
	cur := h.root.Sibling

	for cur != nil {
		if min.Key > cur.Key {
			minPrev = prev
			min = cur
		}
		prev = cur
		cur = cur.Sibling
	}

	minValue := min.Key
	if minPrev == nil {
		h.root = min.Sibling
	} else {
		minPrev.Sibling = min.Sibling
	}

	root := min.Child
	child := min.Child
	for child != nil {
		child.Parent = nil
		child = child.Sibling
	}

	h.root = merge(h.root, root)

	return minValue
}

func (h BinominalHeap) Change(idx, val int) {
	//TODO implement me
	panic("implement me")
}

func (h BinominalHeap) Size() int {
	sz := 0
	cur := h.root
	for cur != nil {
		sz += 1 << cur.Degree
		cur = cur.Sibling
	}

	return sz
}

func merge(h1, h2 *BinominalNode) *BinominalNode {
	cur1 := h1
	cur2 := h2

	result := BinominalNode{}
	curResult := &result

	for cur1 != nil && cur2 != nil {
		if cur1.Degree < cur2.Degree {
			curResult.Sibling = cur1
			curResult = curResult.Sibling
			cur1 = cur1.Sibling
		} else {
			curResult.Sibling = cur2
			curResult = curResult.Sibling
			cur2 = cur2.Sibling
		}
	}

	for cur1 != nil {
		curResult.Sibling = cur1
		curResult = curResult.Sibling
		cur1 = cur1.Sibling
	}

	for cur2 != nil {
		curResult.Sibling = cur2
		curResult = curResult.Sibling
		cur2 = cur2.Sibling
	}

	curResult = result.Sibling

	for curResult != nil {
		next := curResult.Sibling
		if next != nil && curResult.Degree == next.Degree {
			if curResult.Key > next.Key {
				curResult.Degree++
				next.Parent = curResult
				next.Sibling = curResult.Child
				curResult.Child = next
			} else {
				next.Degree++
				curResult.Parent = next
				curResult.Sibling = next.Child
				next.Child = curResult
			}

			curResult = next.Sibling
		} else {
			curResult = next
		}
	}

	return result.Sibling
}
