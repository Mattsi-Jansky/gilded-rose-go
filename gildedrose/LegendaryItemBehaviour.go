package gildedrose

type LegendaryItem struct {
	item *Item
}

func (wrapper *LegendaryItem) UpdateQualityAfterSellIn()  {}
func (wrapper *LegendaryItem) UpdateSellIn()              {}
func (wrapper *LegendaryItem) UpdateQualityBeforeSellIn() {}
