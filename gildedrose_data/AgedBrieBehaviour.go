package gildedrose_data

func agedBrieUpdateSellIn(item *Item) {
	item.SellIn = item.SellIn - 1
}

func agedBrieUpdateQualityAfterSellIn(item *Item) {
	if item.SellIn < 0 && item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func agedBrieUpdateQualityBeforeSellIn(item *Item) {
	if item.Quality < 50 {
		item.Quality += 1
	}
}

func agedBrieSelector(item *Item) bool {
	return item.Name == "Aged Brie"
}

var AgedBrieBehaviour = Behaviour{Select: agedBrieSelector, UpdateQualityBeforeSellIn: agedBrieUpdateQualityBeforeSellIn, UpdateSellIn: agedBrieUpdateSellIn, UpdateQualityAfterSellIn: agedBrieUpdateQualityAfterSellIn}
