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

func TestShouldDegradeQualityByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foobar", 5, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 4
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestShouldReduceSellInByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foobar", 5, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 4
	if items[0].SellIn != expected {
		t.Errorf("SellIn: Expected %d but got %d ", expected, items[0].SellIn)
	}
}
