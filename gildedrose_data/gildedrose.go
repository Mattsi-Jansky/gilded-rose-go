package gildedrose_data

var behaviours = []Behaviour{LegendaryBehaviour, AgedBrieBehaviour, BackstagePassBehaviour, RegularBehaviour}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		for _, behaviour := range behaviours {
			if behaviour.Select(item) {
				behaviour.UpdateQualityBeforeSellIn(item)
				behaviour.UpdateSellIn(item)
				behaviour.UpdateQualityAfterSellIn(item)
				break
			}
		}
	}

}
