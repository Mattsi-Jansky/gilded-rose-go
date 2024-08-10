package gildedrose

type ItemBehaviour interface {
	UpdateQualityBeforeSellIn()
	UpdateSellIn()
	UpdateQualityAfterSellIn()
}
