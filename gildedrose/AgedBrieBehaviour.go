package gildedrose

type AgedBrieItem struct {
	item *Item
}

func (wrapper *AgedBrieItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *AgedBrieItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 && wrapper.item.Quality < 50 {
		wrapper.item.Quality = wrapper.item.Quality + 1
	}
}

func (wrapper *AgedBrieItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality < 50 {
		wrapper.item.Quality += 1
	}
}
