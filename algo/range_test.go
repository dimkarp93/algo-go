package algo

import (
	"algo/util"
	"errors"
	"math"
	"testing"
)

var benchRangeTestData []RangeTestData

type RangeQuery struct {
	left  int
	right int
}

type RangeTestData struct {
	source  []int
	queries []RangeQuery
}

func TestSqrtSum(t *testing.T) {
	data := []int{2, 3, 5, 7, 6, 4}
	sqrt := InitSqrt(data, func(a, b int) int { return a + b })

	doRangeTest(sqrt, -1, 2, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 1, 12, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 3, 2, 0, errors.New("no correct range: left index: 3 more then right index: 2"), t)
	doRangeTest(sqrt, 0, 5, 27, nil, t)
	doRangeTest(sqrt, 0, 2, 10, nil, t)
	doRangeTest(sqrt, 3, 5, 17, nil, t)
	doRangeTest(sqrt, 1, 4, 21, nil, t)
	doRangeTest(sqrt, 2, 3, 12, nil, t)
	doRangeTest(sqrt, 2, 2, 5, nil, t)
	doRangeTest(sqrt, 3, 3, 7, nil, t)
}

func TestSqrtMin(t *testing.T) {
	data := []int{2, 3, 5, 7, 6, 4}
	sqrt := InitSqrt(data, func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	})

	doRangeTest(sqrt, -1, 2, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 1, 12, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 3, 2, 0, errors.New("no correct range: left index: 3 more then right index: 2"), t)
	doRangeTest(sqrt, 0, 5, 2, nil, t)
	doRangeTest(sqrt, 0, 2, 2, nil, t)
	doRangeTest(sqrt, 3, 5, 4, nil, t)
	doRangeTest(sqrt, 1, 4, 3, nil, t)
	doRangeTest(sqrt, 2, 3, 5, nil, t)
	doRangeTest(sqrt, 2, 2, 5, nil, t)
	doRangeTest(sqrt, 3, 3, 7, nil, t)
}

func TestPrefixSum(t *testing.T) {
	data := []int{2, 3, 5, 7, 6, 4}
	sqrt := InitPrefix(data, func(a, b int) int { return a + b }, func(a, b int) int { return a - b })

	doRangeTest(sqrt, -1, 2, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 1, 12, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 3, 2, 0, errors.New("no correct range: left index: 3 more then right index: 2"), t)
	doRangeTest(sqrt, 0, 5, 27, nil, t)
	doRangeTest(sqrt, 0, 2, 10, nil, t)
	doRangeTest(sqrt, 3, 5, 17, nil, t)
	doRangeTest(sqrt, 1, 4, 21, nil, t)
	doRangeTest(sqrt, 2, 3, 12, nil, t)
	doRangeTest(sqrt, 2, 2, 5, nil, t)
	doRangeTest(sqrt, 3, 3, 7, nil, t)
}

func TestSegmentationTree(t *testing.T) {
	data := []int{2, 3, 5, 7, 6, 4}
	sqrt := InitSegmentation(data, func(a, b int) int { return a + b })

	doRangeTest(sqrt, -1, 2, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 1, 12, 0, errors.New("index out of range"), t)
	doRangeTest(sqrt, 3, 2, 0, errors.New("no correct range: left index: 3 more then right index: 2"), t)
	doRangeTest(sqrt, 0, 5, 27, nil, t)
	doRangeTest(sqrt, 0, 2, 10, nil, t)
	doRangeTest(sqrt, 3, 5, 17, nil, t)
	doRangeTest(sqrt, 1, 4, 21, nil, t)
	doRangeTest(sqrt, 2, 3, 12, nil, t)
	doRangeTest(sqrt, 2, 2, 5, nil, t)
	doRangeTest(sqrt, 3, 3, 7, nil, t)
}

func BenchmarkPrefix(b *testing.B) {
	doBenchmark(func(source []int) Range {
		return InitPrefix(
			source,
			func(a, b int) int { return a + b },
			func(a, b int) int { return a - b },
		)
	})
}

func BenchmarkSqrtDecomposition(b *testing.B) {
	doBenchmark(func(source []int) Range {
		return InitSqrt(
			source,
			func(a, b int) int { return a + b },
		)
	})
}

func BenchmarkSegmentation(b *testing.B) {
	doBenchmark(func(source []int) Range {
		return InitSegmentation(
			source,
			func(a, b int) int { return a + b },
		)
	})
}

func doRangeTest(rng Range, a, b, expected int, expErr error, t *testing.T) {
	actual, actErr := rng.GetRange(a, b)

	if actErr != nil && expErr != nil {
		if actErr.Error() != expErr.Error() {
			t.Errorf("errors not equals: expected_error: '%v', actual_error: '%v'", expErr, actErr)
		}
	}

	if actErr != nil && expErr == nil {
		t.Errorf("errors: '%v' was caused, but no expected", actErr)
	}

	if actErr == nil && expErr != nil {
		t.Errorf("erros: '%v' was expected, but not caused", expErr)
	}

	if actual != expected {
		t.Errorf("values not equals: expected: '%v', actual: '%v'", expected, actual)
	}
}

func doBenchmark(makeRange func([]int) Range) {
	for _, testData := range genBenchRangeTests() {
		prefix := makeRange(testData.source)

		for _, q := range testData.queries {
			prefix.GetRange(q.left, q.right)
		}
	}
}

func genBenchRangeTests() []RangeTestData {
	if len(benchData) == 0 {
		benchRangeTestData = genRangeTests(1_000, 1_000, 10_000, 100, 1_000)
	}

	return benchRangeTestData
}

func genRangeTests(n int, minSourceSize, maxSourceSize int, minQueryCount, maxQueryCount int) []RangeTestData {
	result := make([]RangeTestData, n)

	for i := 0; i < n; i++ {
		result[i] = genRangeTest(minSourceSize, maxSourceSize, minQueryCount, maxQueryCount)
	}

	return result
}

func genRangeTest(minSourceSize, maxSourceSize int, minQueryCount, maxQueryCount int) RangeTestData {
	n := util.RandomRange(minSourceSize, maxSourceSize)
	source := util.GenIntSeries(n, math.MinInt, math.MaxInt)

	queryCount := util.RandomRange(minQueryCount, maxQueryCount)
	leftQueries := util.GenIntSeries(queryCount, 0, n-1)
	rightQueries := util.GenIntSeries(queryCount, 0, n-1)

	queries := make([]RangeQuery, queryCount)

	for i := 0; i < n; i++ {
		queries[i] = RangeQuery{left: leftQueries[i], right: rightQueries[i]}
	}

	return RangeTestData{
		source:  source,
		queries: queries,
	}
}
