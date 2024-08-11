package gildedrose_data_test

import (
	"testing"

	gildedrose "github.com/emilybache/gildedrose-refactoring-kata/gildedrose_data"
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

func TestShouldDegradeQualityTwiceAfterSellDate(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foobar", 0, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 3
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestQualityShouldNotGoBelow0(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foobar", 5, 0},
	}

	gildedrose.UpdateQuality(items)

	const expected = 0
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestSellInShouldGoBelow0(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foobar", 0, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = -1
	if items[0].SellIn != expected {
		t.Errorf("SellIn: Expected %d but got %d ", expected, items[0].SellIn)
	}
}

func TestAgeBrieShouldUpgradeQualityByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 5, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 6
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestAgeBrieShouldUpgradeQualityByOneEvenPastSellDate(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", -1, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 7
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestAgeBrieShouldNotGoOver50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 5, 50},
	}

	gildedrose.UpdateQuality(items)

	const expected = 50
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestSulfurasDoesNotChangeSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 5, 50},
	}

	gildedrose.UpdateQuality(items)

	const expected = 5
	if items[0].SellIn != expected {
		t.Errorf("SellIn: Expected %d but got %d ", expected, items[0].SellIn)
	}
}

func TestSulfurasDoesNotChangeQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 5, 50},
	}

	gildedrose.UpdateQuality(items)

	const expected = 50
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestBackstagePassesIncreaseOnceQualityOver10Day(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 11, 1},
	}

	gildedrose.UpdateQuality(items)

	const expected = 2
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestBackstagePassesIncreaseTwiceQuality10Day(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 10, 1},
	}

	gildedrose.UpdateQuality(items)

	const expected = 3
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestBackstagePassesIncreaseThriceQuality5DayOrBelow(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 5, 1},
	}

	gildedrose.UpdateQuality(items)

	const expected = 4
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}

func TestBackstagePassesQualityTo0AfterSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 0, 5},
	}

	gildedrose.UpdateQuality(items)

	const expected = 0
	if items[0].Quality != expected {
		t.Errorf("Quality: Expected %d but got %d ", expected, items[0].Quality)
	}
}
