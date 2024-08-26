// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"math/rand"
	"sort"
)

// 发牌
type Dealer struct {
	EachHandCount int
	Pool          []Card //牌池
}

func NewDealer() *Dealer {
	d := &Dealer{
		Pool:          make([]Card, MaxNCard),
		EachHandCount: 17,
	}
	for i := 0; i < MaxNCard; i++ {
		d.Pool[i] = Card(i + 1)
	}
	return d
}

// 洗牌
func (d *Dealer) Shuffle() {
	rand.Shuffle(len(d.Pool), func(i, j int) {
		d.Pool[i], d.Pool[j] = d.Pool[j], d.Pool[i]
	})
}

// 发牌
func (d *Dealer) Deal() Cards {
	if len(d.Pool) == 0 {
		return nil
	}
	var cnt = d.EachHandCount
	if len(d.Pool) < cnt {
		cnt = len(d.Pool)
	}
	var cards = make(Cards, cnt)
	copy(cards, d.Pool[:cnt])
	d.Pool = d.Pool[cnt:]
	sort.Sort(cards)
	return cards
}
