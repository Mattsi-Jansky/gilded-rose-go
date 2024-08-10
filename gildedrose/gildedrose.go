package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const TypeAgedBrie = "Aged Brie"
const TypePasses = "Backstage passes to a TAFKAL80ETC concert"
const TypeLegendary = "Sulfuras, Hand of Ragnaros"

type RegularItem struct {
	item *Item
}

type LegendaryItem struct {
	item *Item
}

type AgedBrieItem struct {
	item *Item
}

type BackStagePassItem struct {
	item *Item
}

func (wrapper *RegularItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *LegendaryItem) UpdateSellIn() {}

func (wrapper *AgedBrieItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *BackStagePassItem) UpdateSellIn() {
	wrapper.item.SellIn = wrapper.item.SellIn - 1
}

func (wrapper *RegularItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 && wrapper.item.Quality > 0 {
		wrapper.item.Quality = wrapper.item.Quality - 1
	}
}

func (wrapper *LegendaryItem) UpdateQualityAfterSellIn() {}

func (wrapper *AgedBrieItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 && wrapper.item.Quality < 50 {
		wrapper.item.Quality = wrapper.item.Quality + 1
	}
}

func (wrapper *BackStagePassItem) UpdateQualityAfterSellIn() {
	if wrapper.item.SellIn < 0 {
		wrapper.item.Quality = 0
	}
}

func (wrapper *RegularItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality > 0 {
		wrapper.item.Quality -= 1
	}
}

func (wrapper *LegendaryItem) UpdateQualityBeforeSellIn() {}

func (wrapper *AgedBrieItem) UpdateQualityBeforeSellIn() {
	if wrapper.item.Quality < 50 {
		wrapper.item.Quality += 1
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

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		var behaviour ItemBehaviour = nil
		switch items[i].Name {
		case TypeLegendary:
			tmp := new(LegendaryItem)
			tmp.item = items[i]
			behaviour = tmp
		case TypeAgedBrie:
			tmp := new(AgedBrieItem)
			tmp.item = items[i]
			behaviour = tmp
		case TypePasses:
			tmp := new(BackStagePassItem)
			tmp.item = items[i]
			behaviour = tmp
		default:
			tmp := new(RegularItem)
			tmp.item = items[i]
			behaviour = tmp
		}

		behaviour.UpdateQualityBeforeSellIn()
		behaviour.UpdateSellIn()
		behaviour.UpdateQualityAfterSellIn()
	}

}
