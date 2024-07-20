package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func LowerQualityBy1(item *Item) {
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
}

func RaiseQualityBy1(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Name != "Sulfuras, Hand of Ragnaros" {
				LowerQualityBy1(items[i])
			}
		} else {
			RaiseQualityBy1(items[i])
			if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
				if items[i].SellIn < 11 {
					RaiseQualityBy1(items[i])
				}
				if items[i].SellIn < 6 {
					RaiseQualityBy1(items[i])
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Name != "Sulfuras, Hand of Ragnaros" {
						LowerQualityBy1(items[i])
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				RaiseQualityBy1(items[i])
			}
		}
	}
}
