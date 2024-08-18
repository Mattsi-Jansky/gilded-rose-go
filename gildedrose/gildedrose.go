package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const TypeAgedBrie = "Aged Brie"
const TypePasses = "Backstage passes to a TAFKAL80ETC concert"
const TypeLegendary = "Sulfuras, Hand of Ragnaros"
const TypeConjured = "Conjured Mana Cake"

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
		case TypeConjured:
			tmp := new(ConjuredItem)
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
