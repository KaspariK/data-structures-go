package main

import (
	"fmt"
	"reflect"
	"testing"
)

var ll LinkedList

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		items []interface{}
		want []interface{}
	}{
		"Nil list": {items: []interface{}{1,2,3}, want: []interface{}{1,2,3}},
	}

	for name, tc := range tests {
		fmt.Println(tc)
		for i := range tc.items {
			ll.Append(i)
		}

		got := ll.Contents()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}