package algo

import (
	"algo/util"
	"errors"
	"fmt"
	"math"
)

type Range interface {
	GetRange(int, int) (int, error)
	GetLength() int
}

type SqrtDecomposition struct {
	batch  []int
	source []int
	f      func(int, int) int
	n      int
	sz     int
}

type SegmentationTree struct {
	tree   []int
	source []int
	f      func(int, int) int
}

type Prefix struct {
	prefixes []int
	source   []int
	f        func(int, int) int
	g        func(int, int) int
	n        int
}

func ValidateRange(r Range, a, b int) error {
	if a > b {
		return fmt.Errorf("no correct range: left index: %v more then right index: %v", a, b)
	}

	if a >= r.GetLength() || b >= r.GetLength() || a < 0 || b < 0 {
		return errors.New("index out of range")
	}

	return nil
}

func InitSqrt(source []int, f func(int, int) int) SqrtDecomposition {
	n := len(source)
	if n == 0 {
		return SqrtDecomposition{}
	}

	count := int(math.Ceil(math.Sqrt(float64(n))))
	size := count

	batch := make([]int, size)

	for i := 0; i < count; i++ {
		if i*size >= n {
			break
		}
		res := source[i*size]
		for j := 1; j < size; j++ {
			if j+i*size >= n {
				break
			}
			res = f(res, source[j+i*size])
		}
		batch[i] = res
	}

	return SqrtDecomposition{
		batch:  batch,
		source: source,
		f:      f,
		n:      n,
		sz:     size,
	}
}

func (s SqrtDecomposition) GetLength() int {
	return s.n
}

func (s SqrtDecomposition) GetRange(a, b int) (int, error) {
	err := ValidateRange(s, a, b)
	if err != nil {
		return 0, err
	}

	if a == b {
		return s.source[a], nil
	}

	firstBatch := a / s.sz
	if a%s.sz > 0 {
		firstBatch++
	}
	lastBatch := b / s.sz
	l := firstBatch * s.sz
	r := (lastBatch + 1) * s.sz
	if r >= s.n {
		r = s.n - 1
	}

	if firstBatch == lastBatch {
		res := s.source[a]
		for i := a + 1; i <= b; i++ {
			res = s.f(res, s.source[i])
		}
		return res, nil
	}

	hasLeftPart := false
	leftPart := 0
	if a < l {
		hasLeftPart = true

		leftPart = s.source[a]
		for i := a + 1; i < l; i++ {
			leftPart = s.f(leftPart, s.source[i])
		}
	}

	middlePart := s.batch[firstBatch]
	for batch := firstBatch + 1; batch <= lastBatch; batch++ {
		middlePart = s.f(middlePart, s.batch[batch])
	}

	hasRightPart := false
	rightPart := 0
	if b > r {
		hasRightPart = true

		rightPart = s.source[r]
		for i := r + 1; i <= b; i++ {
			rightPart = s.f(rightPart, s.source[i])
		}
	}

	result := middlePart

	if hasLeftPart {
		result = s.f(leftPart, result)
	}

	if hasRightPart {
		result = s.f(rightPart, result)
	}

	return result, nil
}

func InitPrefix(src []int, direct func(int, int) int, inverse func(int, int) int) Prefix {
	n := len(src)
	if n == 0 {
		return Prefix{}
	}

	prefixes := make([]int, n)
	result := src[0]
	prefixes[0] = result
	for i := 1; i < n; i++ {
		result = direct(result, src[i])
		prefixes[i] = result
	}

	return Prefix{
		source:   src,
		prefixes: prefixes,
		f:        direct,
		g:        inverse,
		n:        n,
	}
}

func (p Prefix) GetLength() int {
	return p.n
}

func (p Prefix) GetRange(a, b int) (int, error) {
	err := ValidateRange(p, a, b)
	if err != nil {
		return 0, err
	}

	if a == 0 {
		return p.prefixes[b], nil
	}

	return p.g(p.prefixes[b], p.prefixes[a-1]), nil
}

func InitSegmentation(source []int, f func(int, int) int) SegmentationTree {
	n := len(source)
	treeSize := 4 * util.FindMinimumOfPowerOf2ThanMore(n)
	tree := make([]int, treeSize)
	mask := make([]bool, treeSize)

	for i := 0; i < treeSize; i++ {
		mask[i] = false
	}

	calcTree(tree, source, mask, 0, 0, n-1, f)

	return SegmentationTree{
		tree:   tree,
		source: source,
		f:      f,
	}
}

func calcTree(tree, source []int, mask []bool, i, l, r int, f func(int, int) int) int {
	if l == r {
		tree[i] = source[l]
		mask[i] = true
		return tree[i]
	}

	if mask[i] {
		tree[i] = f(tree[2*i+1], tree[2*i+2])
		mask[i] = true
		return tree[i]
	}

	mid := (l + r) / 2

	tree[i] = f(
		calcTree(tree, source, mask, 2*i+1, l, mid, f),
		calcTree(tree, source, mask, 2*i+2, mid+1, r, f),
	)

	mask[i] = true
	return tree[i]
}

func (t SegmentationTree) GetLength() int {
	return len(t.source)
}

func (t SegmentationTree) GetRange(ql, qr int) (int, error) {
	err := ValidateRange(t, ql, qr)
	if err != nil {
		return 0, err
	}

	return t.getRange(0, 0, len(t.source)-1, ql, qr), nil
}

func (t SegmentationTree) getRange(i, l, r, ql, qr int) int {
	if l == ql && r == qr {
		return t.tree[i]
	}

	mid := (l + r) / 2

	if qr <= mid {
		return t.getRange(2*i+1, l, mid, ql, qr)
	} else if ql > mid {
		return t.getRange(2*i+2, mid+1, r, ql, qr)
	} else {
		return t.f(t.getRange(2*i+1, l, mid, ql, mid), t.getRange(2*i+2, mid+1, r, mid+1, qr))
	}
}
