package gildedrose_data

type Behaviour struct {
	Select                    func(*Item) bool
	UpdateQualityBeforeSellIn func(*Item)
	UpdateSellIn              func(*Item)
	UpdateQualityAfterSellIn  func(*Item)
}
