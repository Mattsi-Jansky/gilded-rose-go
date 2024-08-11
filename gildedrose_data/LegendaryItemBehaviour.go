package gildedrose_data

func legendaryUpdateQualityAfterSellIn(_ *Item)  {}
func legendaryUpdateSellIn(_ *Item)              {}
func legendaryUpdateQualityBeforeSellIn(_ *Item) {}
func legendarySelector(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

var LegendaryBehaviour = Behaviour{Select: legendarySelector, UpdateQualityBeforeSellIn: legendaryUpdateQualityBeforeSellIn, UpdateSellIn: legendaryUpdateSellIn, UpdateQualityAfterSellIn: legendaryUpdateQualityAfterSellIn}
