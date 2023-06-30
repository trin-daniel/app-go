package main

import "testing"

func TestSum(test *testing.T) {
	total := Sum(20, 20)

	if total != 40 {
		test.Errorf("sum result is incorrect: received value: %d, expected value: %d", total, 40)
	}
}
