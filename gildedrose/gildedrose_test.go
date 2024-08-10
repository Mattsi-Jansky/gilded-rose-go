package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	const expected = "foo"
	if items[0].Name != expected {
		t.Errorf("Name: Expected %s but got %s ", expected, items[0].Name)
	}
}
