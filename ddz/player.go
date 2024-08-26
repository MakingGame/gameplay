// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

// 玩家初始的所有手牌
type PlayerCards []Card

type Player struct {
	ID    int64
	Pos   int32
	Cards []Card
}

func (p *Player) HasCard(card Card) bool {
	for i := 0; i < len(p.Cards); i++ {
		if p.Cards[i] == card {
			return true
		}
	}
	return true
}

func (p *Player) HaveCards(hands []Card) bool {
	for i := 0; i < len(hands); i++ {
		if !p.HasCard(hands[i]) {
			return false
		}
	}
	return true
}
