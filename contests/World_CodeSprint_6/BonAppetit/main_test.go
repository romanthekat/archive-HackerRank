package main

import "testing"

func TestOvercharge(t *testing.T) {
	input := Input{
		itemsCount:4,
		allergicIndex:1,
		items: []int{3, 10, 2, 9},
		charged:12,
	}

	output := analyzeInput(input)

	overcharge := output.overcharge
	if overcharge != 5 {
		t.Error("overcharge value must be 5, but equals:", overcharge)
	}
}

func TestBonAppetit(t *testing.T) {
	input := Input{
		itemsCount:4,
		allergicIndex:1,
		items: []int{3, 10, 2, 9},
		charged:7,
	}

	output := analyzeInput(input)

	overcharge := output.overcharge
	if overcharge != 0 {
		t.Error("overcharge value must be 0, but equals:", overcharge)
	}
}
