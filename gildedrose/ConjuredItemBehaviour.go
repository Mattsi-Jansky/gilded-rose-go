package gildedrose

type ConjuredItem struct {
	item *Item
}

func (wrapper *ConjuredItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 && wrapper.item.Quality > 0 {
		wrapper.item.Quality = wrapper.item.Quality - 1
	}
}
func (wrapper *ConjuredItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}
func (wrapper *ConjuredItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality > 0 {
		wrapper.item.Quality -= 2
	}
}
