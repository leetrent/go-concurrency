package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	for ii := 0; ii < 10; ii++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()

		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(orderFinished))
		}
	}
}
