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
