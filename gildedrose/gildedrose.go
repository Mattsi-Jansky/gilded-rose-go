package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const TypeAgedBrie = "Aged Brie"
const TypePasses = "Backstage passes to a TAFKAL80ETC concert"
const TypeLegendary = "Sulfuras, Hand of Ragnaros"

type ItemBehaviour interface {
	UpdateQualityBeforeSellIn()
	UpdateSellIn()
	UpdateQualityAfterSellIn()
}

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

func (item *RegularItem) UpdateSellIn() {
	item.item.SellIn = item.item.SellIn - 1
}

func (item *LegendaryItem) UpdateSellIn() {}

func (item *AgedBrieItem) UpdateSellIn() {
	item.item.SellIn = item.item.SellIn - 1
}

func (item *BackStagePassItem) UpdateSellIn() {
	item.item.SellIn = item.item.SellIn - 1
}

func (item *RegularItem) UpdateQualityAfterSellIn() {
	if item.item.SellIn < 0 && item.item.Quality > 0 {
		item.item.Quality = item.item.Quality - 1
	}
}

func (item *LegendaryItem) UpdateQualityAfterSellIn() {}

func (item *AgedBrieItem) UpdateQualityAfterSellIn() {
	if item.item.SellIn < 0 && item.item.Quality < 50 {
		item.item.Quality = item.item.Quality + 1
	}
}

func (item *BackStagePassItem) UpdateQualityAfterSellIn() {
	if item.item.SellIn < 0 {
		item.item.Quality = 0
	}
}

func (item *RegularItem) UpdateQualityBeforeSellIn() {
	if item.item.Quality > 0 {
		item.item.Quality -= 1
	}
}

func (item *LegendaryItem) UpdateQualityBeforeSellIn() {}

func (item *AgedBrieItem) UpdateQualityBeforeSellIn() {
	if item.item.Quality < 50 {
		item.item.Quality += 1
	}
}

func (item *BackStagePassItem) UpdateQualityBeforeSellIn() {
	if item.item.Quality < 50 {

		item.item.Quality += 1
		if item.item.SellIn < 11 {
			if item.item.Quality < 50 {
				item.item.Quality = item.item.Quality + 1
			}
		}
		if item.item.SellIn < 6 {
			if item.item.Quality < 50 {
				item.item.Quality = item.item.Quality + 1
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
