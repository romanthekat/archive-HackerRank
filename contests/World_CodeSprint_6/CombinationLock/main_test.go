package main

import "testing"

func TestCombinationLock(t *testing.T) {
	lines := []string{"1 2 9 5 7", "1 3 2 0 7"}
	moves := calculateMovesNumber(lines)

	if moves != 9 {
		t.Error("Moves count must be 9, but equals to ", moves)
	}
}