package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const TypeAgedBrie = "Aged Brie"
const TypePasses = "Backstage passes to a TAFKAL80ETC concert"
const TypeLegendary = "Sulfuras, Hand of Ragnaros"

func (item *Item) UpdateItemQuality(behaviour ItemBehaviour) {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality > 0 {
			behaviour.DecrementQuality()
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > 0 {
					behaviour.DecrementQuality()
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
	DecrementQuality()
}

type RegularItem struct {
	item *Item
}

func (item *RegularItem) DecrementQuality() {
	item.item.Quality -= 1
}

type LegendaryItem struct {
	item *Item
}

func (item *LegendaryItem) DecrementQuality() {}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		var behaviour ItemBehaviour = nil
		switch items[i].Name {
		case TypeLegendary:
			tmp := new(LegendaryItem)
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
