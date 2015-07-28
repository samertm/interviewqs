package interviewcake

import (
	"sort"
	"testing"
)

func TestCondenseMeetingTimes(t *testing.T) {
	testcases := []struct {
		input  []Meeting
		output []Meeting
	}{{
		input:  []Meeting{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}},
		output: []Meeting{{0, 1}, {3, 8}, {9, 12}},
	}, {
		input:  []Meeting{{1, 2}, {2, 3}},
		output: []Meeting{{1, 3}},
	}, {
		input:  []Meeting{{1, 5}, {2, 3}},
		output: []Meeting{{1, 5}},
	}, {
		input:  []Meeting{{1, 10}, {2, 6}, {3, 5}, {7, 9}},
		output: []Meeting{{1, 10}},
	}}
	meetingsEqual := func(m0, m1 []Meeting) bool {
		if len(m0) != len(m1) {
			return false
		}
		sort.Sort(SortableMeetings(m0))
		sort.Sort(SortableMeetings(m1))
		for i := range m0 {
			if m0[i] != m1[i] {
				return false
			}
		}
		return true
	}
	for _, test := range testcases {
		if got := CondenseMeetingTimes(test.input); !meetingsEqual(got, test.output) {
			t.Errorf("Got %+v, wanted %+v", got, test.output)
		}
	}

}
