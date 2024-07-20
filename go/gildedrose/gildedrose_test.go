package gildedrose_test

import (
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

func TestLowerQualityBy1(t *testing.T) {
	tests := []struct {
		name            string
		initialQuality  int
		expectedQuality int
	}{
		{"Quality decreases by 1", 10, 9},
		{"Quality at 0 stays 0", 0, 0},
		{"High initial quality decreases by 1", 50, 49},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: "foo", Quality: test.initialQuality}
			gildedrose.LowerQualityBy1(item)
			if item.Quality != test.expectedQuality {
				t.Errorf("%s: Expected quality %d but got %d", test.name, test.expectedQuality, item.Quality)
			}
		})
	}
}
