package gildedrose_data

func backstagePassUpdateSellIn(item *Item) {
	item.SellIn = item.SellIn - 1
}

func backstagePassUpdateQualityAfterSellIn(item *Item) {
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

func backstagePassUpdateQualityBeforeSellIn(item *Item) {
	if item.Quality < 50 {

		item.Quality += 1
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}

func backstagePassSelector(item *Item) bool {
	return item.Name == "Backstage passes to a TAFKAL80ETC concert"
}

var BackstagePassBehaviour = Behaviour{Select: backstagePassSelector, UpdateQualityBeforeSellIn: backstagePassUpdateQualityBeforeSellIn, UpdateSellIn: backstagePassUpdateSellIn, UpdateQualityAfterSellIn: backstagePassUpdateQualityAfterSellIn}
