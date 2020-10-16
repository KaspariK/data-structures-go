package list

import (
	"reflect"
	"testing"
)

// TODO: More test cases
func TestPrepend(t *testing.T) {
	tests := map[string]struct {
		items []interface{}
		want []interface{}
	}{
		"Single item": {items:[]interface{}{1}, want:[]interface{}{1}},
		"Multiple items": {items: []interface{}{1,2,3}, want: []interface{}{3,2,1}},
		"Single string": {items: []interface{}{"Kevin"}, want: []interface{}{"Kevin"}},
		"Multiple strings": {items: []interface{}{"The", "quick", "brown", "fox"}, want: []interface{}{"fox", "brown", "quick", "The"}},
	}

	for name, tc := range tests {
		var ll LinkedList

		for _, item := range tc.items {
			ll.Prepend(item)
		}

		got := ll.Contents()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

// TODO: More test cases
func TestAppend(t *testing.T) {
	tests := map[string]struct {
		items []interface{}
		want []interface{}
	}{
		"Single int": {items: []interface{}{1}, want: []interface{}{1}},
		"Multiple ints": {items: []interface{}{1,2,3}, want: []interface{}{1,2,3}},
		"Single string": {items: []interface{}{"Kevin"}, want: []interface{}{"Kevin"}},
		"Multiple strings": {items: []interface{}{"The", "quick", "brown", "fox"}, want: []interface{}{"The", "quick", "brown", "fox"}},
	}

	for name, tc := range tests {
		var ll LinkedList

		for _, item := range tc.items {
			ll.Append(item)
		}

		got := ll.Contents()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestClear(t *testing.T) {
	tests := map[string]struct{
		items []interface{}
	}{
		"No items": { items: []interface{}{}},
		"Single item": { items: []interface{}{1}},
		"Multiple items": { items: []interface{}{1, 2, 3}},
	}

	for name, tc := range tests {
		var ll LinkedList

		for _, item := range tc.items {
			ll.Append(item)
		}

		ll.Clear()
		got := ll.Contents()
		if !ll.IsEmpty() {
			t.Errorf("%s: expected: [], got: %v", name, got)
		}
	}
}
func TestSize(t *testing.T) {
	tests := map[string]struct{
		items []interface{}
		want int
	} {
		"No items": {items: []interface{}{}, want: 0},
		"One int": {items: []interface{}{1}, want: 1},
		"Multiple items": {items: []interface{}{1, 2, 3}, want: 3},
	}

	for name, tc := range tests {
		var ll LinkedList

		for _, item := range tc.items {
			ll.Append(item)
		}

		got := ll.len
		if got != tc.want {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := map[string]struct{
		items []interface{}
		want bool
	}{
		"Empty": {items: []interface{}{}, want: true},
		"Not empty": {items: []interface{}{"The", "quick", "brown", "fox"}, want: false},
	}

	for name, tc := range tests {
		var ll LinkedList

		for _, item := range tc.items {
			ll.Append(item)
		}

		got := ll.IsEmpty()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
