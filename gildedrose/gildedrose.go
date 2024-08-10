package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const TypeAgedBrie = "Aged Brie"
const TypePasses = "Backstage passes to a TAFKAL80ETC concert"
const TypeLegendary = "Sulfuras, Hand of Ragnaros"

func (item *Item) UpdateItemQuality(behaviour ItemBehaviour) {

	behaviour.UpdateQuality()

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > 0 {
					behaviour.UpdateQuality()
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}

type ItemBehaviour interface {
	UpdateQuality()
}

type RegularItem struct {
	item *Item
}

func (item *RegularItem) UpdateQuality() {
	if item.item.Quality > 0 {
		item.item.Quality -= 1
	}
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

func (item *LegendaryItem) UpdateQuality() {}

func (item *AgedBrieItem) UpdateQuality() {
	if item.item.Quality < 50 {
		item.item.Quality += 1
	}
}

func (item *BackStagePassItem) UpdateQuality() {
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
		items[i].UpdateItemQuality(behaviour)
	}

}
