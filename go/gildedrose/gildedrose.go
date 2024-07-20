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
	item.SellIn = item.SellIn - 1
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Name != "Sulfuras, Hand of Ragnaros" {
				LowerQualityByNumber(items[i], 1)
			}
		} else {
			RaiseQualityByNumber(items[i], 1)
			if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
				if items[i].SellIn < 11 {
					RaiseQualityByNumber(items[i], 1)
				}
				if items[i].SellIn < 6 {
					RaiseQualityByNumber(items[i], 1)
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			UpdateSellIn(items[i])
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Name != "Sulfuras, Hand of Ragnaros" {
						LowerQualityByNumber(items[i], 1)
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				RaiseQualityByNumber(items[i], 1)
			}
		}
	}
}
