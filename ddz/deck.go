// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

type Deck struct {
	dealer  *Dealer
	players []*Player
	pool    []Card
}

func (d *Deck) Play(player *Player, category HandsCategory, cards []Card) {

}

func CheckCardsRange(cards []Card) bool {
	for i := 0; i < len(cards); i++ {
		if !(cards[i] >= MinNCard && cards[i] <= MaxNCard) {
			return false
		}
	}
	return true
}
