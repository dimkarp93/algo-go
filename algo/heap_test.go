package algo

import (
	"algo/util"
	"testing"
)

var binaryHeapTestData []BinaryHeap

func TestBinaryHeapBuild(t *testing.T) {
	doTestHeap(t, BuildBinaryHeap([]int{1, 5, 3, 7, 2, 0, 9, 12, 3, 6, 0}))
}

func TestBinominalHeapBuild(t *testing.T) {
	doTestHeap(t, BuildBinominalHeap([]int{1, 5, 3, 7, 2, 0, 9, 12, 3, 6, 0}))
}

func doTestHeap(t *testing.T, heap Heap) {
	doTestHeapMin(t, 0, heap)
	doTestHeapPop(t, 0, heap)
	doTestHeapPop(t, 0, heap)
	doTestHeapMin(t, 1, heap)
	doTestHeapAdd(t, -5, heap, 5)
	doTestHeapAdd(t, -5, heap, 20)
	doTestHeapAdd(t, -10, heap, -10)
	//doTestHeapChanged(t, -5, heap, 0, 100)
}

func doTestHeapMin(t *testing.T, expected int, heap Heap) {
	sizeOld := heap.Size()
	actual := heap.Min()
	sizeNew := heap.Size()

	if actual != expected {
		t.Errorf("Min incorrect. Expected: '%v', but actual is: '%v'", expected, actual)
	}

	if sizeOld != sizeNew {
		t.Errorf("Min operation has chaged size of heap, but not. "+
			"Old size: '%v', new size: '%v'", sizeOld, sizeNew)
	}
}

func doTestHeapPop(t *testing.T, expected int, heap Heap) {
	sizeOld := heap.Size()
	actual := heap.Pop()
	sizeNew := heap.Size()

	if actual != expected {
		t.Errorf("Min incorrect. Expected: '%v', but actual is: '%v'", expected, actual)
	}

	if sizeOld != sizeNew+1 {
		t.Errorf("New size of heap must be old size minus one, but not. "+
			"Old size: '%v', new size: '%v'", sizeOld, sizeNew)
	}
}

func doTestHeapAdd(t *testing.T, expected int, heap Heap, value int) {
	sizeOld := heap.Size()
	heap.Add(value)
	sizeNew := heap.Size()

	actual := heap.Min()

	if actual != expected {
		t.Errorf("Min incorrect. Expected min of heap: '%v', but actual is: '%v'", expected, actual)
	}

	if sizeOld != sizeNew-1 {
		t.Errorf("New size of heap must be old size plus one, but not. "+
			"Old size: '%v', new size: '%v'", sizeOld, sizeNew)
	}
}

/*
func doTestHeapChanged(t *testing.T, expected int, heap Heap, index, value int) {
	sizeOld := heap.Size()
	heap.Change(index, value)
	sizeNew := heap.Size()

	actual := heap.Min()

	if actual != expected {
		t.Errorf("Min incorrect. Expected min of heap: '%v', but actual is: '%v'", expected, actual)
	}

	if sizeOld != sizeNew {
		t.Errorf("New size of heap must equals old size, but not. "+
			"Old size: '%v', new size: '%v'", sizeOld, sizeNew)
	}
}
*/

func BenchmarkBinaryHeap(b *testing.B) {
	for _, heap := range getBinaryHeapBenchmarkData() {
		doHeapBenchmark(heap)
	}
}

func doHeapBenchmark(heap Heap) {
	heap.Add(util.RandomRange(0, 100_000_000))
	heap.Min()
	heap.Pop()
	heap.Add(util.RandomRange(0, 100_000_000))
	//heap.Change(util.RandomRange(0, heap.Size()-1), util.RandomRange(0, 100_000_000))
	heap.Min()
}

func getBinaryHeapBenchmarkData() []BinaryHeap {
	if len(binaryHeapTestData) == 0 {
		binaryHeapTestData = genBinaryHeap(100, 500, 2_000, 0, 10_000)
	}

	return binaryHeapTestData
}

func genBinaryHeap(n, heapSizeMin, heapSizeMax, valueMin, valueMax int) []BinaryHeap {
	result := make([]BinaryHeap, n)
	for i := 0; i < n; i++ {
		heapSize := util.RandomRange(heapSizeMin, heapSizeMax)
		result[i] = BuildBinaryHeap(util.GenIntSeries(heapSize, valueMin, valueMax))
	}
	return result
}
