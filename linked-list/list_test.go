package list

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		items []interface{}
		want []interface{}
	}{
		"Single item": {items: []interface{}{1}, want: []interface{}{1}},
		"Multiple items": {items: []interface{}{1,2,3}, want: []interface{}{1,2,3}},
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

func TestPrepend(t *testing.T) {
	tests := map[string]struct {
		items []interface{}
		want []interface{}
	}{
		"One item": {items:[]interface{}{1}, want:[]interface{}{1}},
		"Multiple items": {items: []interface{}{1,2,3}, want: []interface{}{3,2,1}},
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