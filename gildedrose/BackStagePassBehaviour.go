package gildedrose

type BackStagePassItem struct {
	item *Item
}

func (wrapper *BackStagePassItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *BackStagePassItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 {
		wrapper.item.Quality = 0
	}
}

func (wrapper *BackStagePassItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality < 50 {

		wrapper.item.Quality += 1
		if wrapper.item.SellIn < 11 {
			if wrapper.item.Quality < 50 {
				wrapper.item.Quality = wrapper.item.Quality + 1
			}
		}
		if wrapper.item.SellIn < 6 {
			if wrapper.item.Quality < 50 {
				wrapper.item.Quality = wrapper.item.Quality + 1
			}
		}
	}
}
