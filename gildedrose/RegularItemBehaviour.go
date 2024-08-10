package gildedrose

type RegularItem struct {
	item *Item
}

func (wrapper *RegularItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *RegularItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 && wrapper.item.Quality > 0 {
		wrapper.item.Quality = wrapper.item.Quality - 1
	}
}

func (wrapper *RegularItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality > 0 {
		wrapper.item.Quality -= 1
	}
}
