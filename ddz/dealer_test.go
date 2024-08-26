// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"testing"
)

func TestDealer_Deal(t *testing.T) {
	var d = NewDealer()
	t.Logf("init: %v", d.Pool)
	d.Shuffle()
	t.Logf("shuffled: %v", d.Pool)
	for i := 0; i < 4; i++ {
		var cards = d.Deal()
		t.Logf("%d-th hands: %v", i+1, cards)
	}
}
