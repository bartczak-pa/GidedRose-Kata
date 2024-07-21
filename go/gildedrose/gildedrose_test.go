package gildedrose_test

import (
	"fmt"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"fixme", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func TestLowerQualityByNumber(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		number          int
		expectedQuality int
	}{
		{"Quality decreases by 1", 10, 1, 9},
		{"Quality decreases by 2", 10, 2, 8},
		{"Quality at 0 stays 0", 0, 1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: "foo", Quality: test.initialQuality}
			gildedrose.LowerQualityByNumber(item, test.number)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}
}

func ExampleLowerQualityByNumber() {
	item := &gildedrose.Item{Name: "foo", Quality: 10}
	gildedrose.LowerQualityByNumber(item, 2)
	fmt.Println(item.Quality)
	// Output: 8
}

func TestRaiseQualityByNumber(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		number          int
		expectedQuality int
	}{
		{"Quality increases by 1", 10, 1, 11},
		{"Quality increases by 2", 10, 2, 12},
		{"Quality at 50 stays 50", 50, 1, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: "foo", Quality: test.initialQuality}
			gildedrose.RaiseQualityByNumber(item, test.number)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}

}

func ExampleRaiseQualityByNumber() {
	item := &gildedrose.Item{Name: "foo", Quality: 10}
	gildedrose.RaiseQualityByNumber(item, 2)
	fmt.Println(item.Quality)
	// Output: 12
}

func TestUpdateSellIn(t *testing.T) {
	tests := []struct {
		name           string
		initialSellIn  int
		expectedSellIn int
		itemName       string
	}{
		{"SellIn decreases by 1", 10, 9, "foo"},
		{"SellIn stays at 0", 0, 0, "Sulfuras, Hand of Ragnaros"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: test.itemName, SellIn: test.initialSellIn}
			gildedrose.UpdateSellIn(item)
			if item.SellIn != test.expectedSellIn {
				t.Errorf("%s: Expected sellIn %d but got %d", test.name, test.expectedSellIn, item.SellIn)
			}
		})

	}
}

func ExampleUpdateSellIn() {
	item := &gildedrose.Item{Name: "foo", SellIn: 10}
	gildedrose.UpdateSellIn(item)
	fmt.Println(item.SellIn)
	// Output: 9

}

func TestUpdateQualityAfterSellIn(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		initialSellIn   int
		expectedQuality int
		itemName        string
	}{
		{"Quality decreases by 1", 10, -1, 9, "foo"},
		{"Quality at 0 stays 0", 0, 0, 0, "foo"},
		{"Quality increases by 1", 10, -1, 11, "Aged Brie"},
		{"Quality at 0 stays 0", 0, 0, 0, "Aged Brie"},
		{"Quality at 0 stays 0", 0, 0, 0, "Sulfuras, Hand of Ragnaros"},
		{"Quality at 0 stays 0", 0, 0, 0, "Backstage passes to a TAFKAL80ETC concert"},
		{"Quality at 0 stays 0", 0, 0, 0, "Conjured Mana Cake"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: test.itemName, Quality: test.initialQuality, SellIn: test.initialSellIn}
			gildedrose.UpdateQualityAfterSellIn(item)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}
}

func ExampleUpdateQualityAfterSellIn() {
	item := &gildedrose.Item{Name: "foo", Quality: 10, SellIn: -1}
	gildedrose.UpdateQualityAfterSellIn(item)
	fmt.Println(item.Quality)
	// Output: 9

}

func TestUpdateBrieQuality(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		expectedQuality int
	}{
		{"Quality increases by 1", 10, 11},
		{"Quality at 50 stays 50", 50, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: "Aged Brie", Quality: test.initialQuality}
			gildedrose.UpdateBrieQuality(item)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}
}

func ExampleUpdateBrieQuality() {
	item := &gildedrose.Item{Name: "Aged Brie", Quality: 10}
	gildedrose.UpdateBrieQuality(item)
	fmt.Println(item.Quality)
	// Output: 11
}

func TestUpdateBackStagePassesQuality(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		initialSellIn   int
		expectedQuality int
	}{
		{"Quality increases by 1", 10, 20, 11},
		{"Quality increases by 2", 10, 10, 12},
		{"Quality increases by 3", 10, 5, 13},
		{"Quality at 50 stays 50", 50, 20, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: test.initialQuality, SellIn: test.initialSellIn}
			gildedrose.UpdateBackStagePassesQuality(item)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}
}

func ExampleUpdateBackStagePassesQuality() {
	item := &gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 10, SellIn: 20}
	gildedrose.UpdateBackStagePassesQuality(item)
	fmt.Println(item.Quality)
	// Output: 11
}
