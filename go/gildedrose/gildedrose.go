package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func LowerQualityByNumber(item *Item, number int) {
	if item.Quality > 0 {
		item.Quality = item.Quality - number
	}
}

func RaiseQualityByNumber(item *Item, number int) {
	if item.Quality < 50 {
		item.Quality = item.Quality + number
	}
}

func UpdateSellIn(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}
}

func UpdateQualityAfterSellIn(item *Item) {
	if item.SellIn < 0 {
		if item.Name != "Sulfuras, Hand of Ragnaros" {
			switch item.Name {
			case "Aged Brie":
				UpdateBrieQuality(item)
			case "Backstage passes to a TAFKAL80ETC concert":
				item.Quality = 0
			default:
				LowerQualityByNumber(item, 1)
			}
		}
	}
}

func UpdateBrieQuality(item *Item) {
	if item.Quality < 50 {
		RaiseQualityByNumber(item, 1)
	}
}

func UpdateBackStagePassesQuality(item *Item) {
	if item.Quality < 50 {
		RaiseQualityByNumber(item, 1)
		if item.SellIn < 11 {
			UpdateBrieQuality(item)
		}
		if item.SellIn < 6 {
			UpdateBrieQuality(item)
		}
	}
}

func UpdateConjuredQuality(item *Item) {
	LowerQualityByNumber(item, 2)
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			switch items[i].Name {
			case "Aged Brie":
				UpdateBrieQuality(items[i])
			case "Backstage passes to a TAFKAL80ETC concert":
				UpdateBackStagePassesQuality(items[i])
			case "Conjured Mana Cake":
				UpdateConjuredQuality(items[i])
			default:
				LowerQualityByNumber(items[i], 1)
			}
		}
		UpdateSellIn(items[i])
		UpdateQualityAfterSellIn(items[i])
	}
}
