package util_test

import (
	"fmt"
	"testing"

	"github.com/jdbann/tilestack/util"
	"gotest.tools/v3/assert"
)

func ExampleIterator() {
	next, done := util.Iterator(2, 0)
	fmt.Println(next(), done())
	fmt.Println(next(), done())
	fmt.Println(next(), done())
	fmt.Println(next(), done())
	// Output: 2 false
	// 1 false
	// 0 true
	// 0 true
}

func TestIterator(t *testing.T) {
	type testCase struct {
		from, to int
		want     []int
	}

	run := func(t *testing.T, tc testCase) {
		next, done := util.Iterator(tc.from, tc.to)
		var got []int
		for !done() {
			got = append(got, next())
		}
		assert.DeepEqual(t, got, tc.want)
		assert.Equal(t, got[len(got)-1], next(), "should continue returning final value")
		assert.Assert(t, done(), "should remain done")
	}

	testCases := []testCase{
		{from: 0, to: 3, want: []int{0, 1, 2, 3}},
		{from: 3, to: 0, want: []int{3, 2, 1, 0}},
		{from: 3, to: -3, want: []int{3, 2, 1, 0, -1, -2, -3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("from: %d, to %d", tc.from, tc.to), func(t *testing.T) {
			run(t, tc)
		})
	}
}
