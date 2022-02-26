package algo

import (
	"algo/util"
	"sort"
	"sync/atomic"
	"testing"
)

type TestData struct {
	Source     []int
	SourceOrig []int
	Expected   []int
}

var manualData = []TestData{
	{
		Source:     []int{3, 2, 1, 4},
		SourceOrig: []int{3, 2, 1, 4},
		Expected:   []int{1, 2, 3, 4},
	},
	{
		Source:     []int{4, 7, 2, 1, 9, 0, 8, 6, 5, 2, 3, 1, 3, 5, 7, 8},
		SourceOrig: []int{4, 7, 2, 1, 9, 0, 8, 6, 5, 2, 3, 1, 3, 5, 7, 8},
		Expected:   []int{0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 7, 8, 8, 9},
	},
}

func genTest(tests int, minSize, maxSize int, minValue, maxValue int) []TestData {
	results := make([]TestData, tests)
	for i := range results {
		results[i] = genTestData(util.RandomRange(minSize, maxSize), minValue, maxValue)
	}
	return results
}

func genTestData(n int, min, max int) TestData {
	source := util.GenIntSeries(n, min, max)
	sourceOrig := util.ArrayCopy(source)
	expected := util.ArrayCopy(source)

	sort.Ints(expected)

	return TestData{
		Source:     source,
		SourceOrig: sourceOrig,
		Expected:   expected,
	}
}

var genData = genTest(50, 100, 10_000, -1000, 1000)
var testData = append(manualData, genData...)

var benchData = genTest(2000, 25_000, 50_000, -100_000, 100_000)
var benchCounter uint64 = 0

func testSort(t *testing.T, sort func([]int) []int) {
	for _, data := range testData {
		source := data.Source
		sourceOrig := data.SourceOrig
		expected := data.Expected

		result := sort(source)

		if !util.ArrEquals(result, expected) {
			t.Errorf("FAILED, result: '%v', expected: '%v'", result, expected)
		}

		if !util.ArrEquals(source, sourceOrig) {
			t.Errorf("FAILED, source: '%v', sourceOrig: '%v'", source, sourceOrig)
		}
	}
}

func wrapSort(sort func([]int)) func([]int) []int {
	return func(arr []int) []int {
		result := util.ArrayCopy(arr)
		sort(result)
		return result
	}
}

func genBenchTestData() []int {
	counter := atomic.AddUint64(&benchCounter, 1)
	return benchData[counter%(uint64)(len(benchData))].Source
}

func TestSelectSort(t *testing.T) {
	testSort(t, SelectSort)
}

func TestInsertSort(t *testing.T) {
	testSort(t, InsertSort)
}

func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort)
}

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort)
}

func TestQuickSort(t *testing.T) {
	testSort(t, QuickSort)
}

func TestHeapSort(t *testing.T) {
	testSort(t, HeapSort)
}

func TestCounterSort(t *testing.T) {
	testSort(t, CounterSort)
}

func TestRadixSort(t *testing.T) {
	testSort(t, RadixSort)
}

func TestInternalSort(t *testing.T) {
	testSort(t, wrapSort(sort.Ints))
}

func BenchmarkSelectSort(b *testing.B) {
	SelectSort(genBenchTestData())
}

func BenchmarkInsertSort(b *testing.B) {
	InsertSort(genBenchTestData())
}

func BenchmarkBubbleSort(b *testing.B) {
	BubbleSort(genBenchTestData())
}

func BenchmarkMergeSort(b *testing.B) {
	MergeSort(genBenchTestData())
}

func BenchmarkQuickSort(b *testing.B) {
	QuickSort(genBenchTestData())
}

func BenchmarkHeapSort(b *testing.B) {
	HeapSort(genBenchTestData())
}

func BenchmarkCounterSort(b *testing.B) {
	CounterSort(genBenchTestData())
}

func BenchmarkRadixSort(b *testing.B) {
	RadixSort(genBenchTestData())
}

func BenchmarkInternalSort(b *testing.B) {
	sortF := wrapSort(sort.Ints)
	sortF(genBenchTestData())
}
