package gildedrose_data

func regularUpdateSellIn(item *Item) {
	item.SellIn = item.SellIn - 1
}

func regularUpdateQualityAfterSellIn(item *Item) {
	if item.SellIn < 0 && item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
}

func regularUpdateQualityBeforeSellIn(item *Item) {
	if item.Quality > 0 {
		item.Quality -= 1
	}
}

func regularSelector(item *Item) bool {
	return true
}

var RegularBehaviour = Behaviour{Select: regularSelector, UpdateQualityBeforeSellIn: regularUpdateQualityBeforeSellIn, UpdateSellIn: regularUpdateSellIn, UpdateQualityAfterSellIn: regularUpdateQualityAfterSellIn}
