// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"strconv"
	"unicode/utf8"
)

// 一张牌，取值为1-54，分别表示:
//
//	【方块3-方块A、方块2】，【梅花3-梅花A、梅花2】，
//	【红桃3-红桃A、红桃2】，【黑桃3-黑桃A、黑桃2】【小王、大王】
type Card int32

func MakeCard(suit CardSuit, point CardPoint) Card {
	switch suit {
	case CardSuitDiamonds:
		return Card(point)
	case CardSuitClubs:
		return Card(EachSuitN + point)
	case CardSuitHearts:
		return Card(EachSuitN*2 + point)
	case CardSuitSpades:
		return Card(EachSuitN*3 + point)
	case CardSuitBlackJoker:
		return CardBlackJoker
	case CardSuitRedJoker:
		return CardRedJoker
	}
	return 0
}

// Card from string, `s`第一个是花色，第二个是牌面值
func CardFromStr(s string) Card {
	switch s {
	case NameBlackJoker:
		return CardBlackJoker // black joker
	case NameRedJoker:
		return CardRedJoker // red joker
	}
	var n = utf8.RuneCountInString(s)
	if !(n == 2 || n == 3) {
		panic("invalid card string: " + s)
	}
	var suit CardSuit
	var pt CardPoint
	r, n := utf8.DecodeRuneInString(s)
	switch string(r) {
	case NameDiamond:
		suit = CardSuitDiamonds
	case NameClub:
		suit = CardSuitClubs
	case NameHeart:
		suit = CardSuitHearts
	case NameSpade:
		suit = CardSuitSpades
	}
	s = s[n:]
	switch s {
	case NameRankA:
		pt = CardPointA
	case NameRankK:
		pt = CardPointK
	case NameRankQ:
		pt = CardPointQ
	case NameRankJ:
		pt = CardPointJ
	case NameRank2:
		pt = CardPoint2
	default:
		// [3 - 10]
		pn, _ := strconv.Atoi(s)
		pt = CardPoint(pn - 2)
	}
	return MakeCard(suit, pt)
}

// 花色
func (c Card) Suit() CardSuit {
	switch c {
	case CardBlackJoker:
		return CardSuitBlackJoker
	case CardRedJoker:
		return CardSuitRedJoker
	}
	if c < MinNCard || c > MaxNCard {
		return SuitNone
	}
	for i := CardSuitDiamonds; i <= CardSuitSpades; i++ {
		if int32(c) > int32(i-1*EachSuitN) && int32(c) <= int32(i*EachSuitN) {
			return i
		}
	}
	return SuitNone
}

func (c Card) SuitString() string {
	switch c.Suit() {
	case CardSuitDiamonds:
		return NameDiamond
	case CardSuitClubs:
		return NameClub
	case CardSuitHearts:
		return NameHeart
	case CardSuitSpades:
		return NameSpade
	case CardSuitBlackJoker:
		return NameBlackJoker
	case CardSuitRedJoker:
		return NameRedJoker
	}
	return "?"
}

// 牌面大小（3-10,J,Q,K,A,2)
func (c Card) Rank() CardPoint {
	switch c {
	case CardBlackJoker:
		return CardPointBJK
	case CardRedJoker:
		return CardPointRJK
	}
	if c < MinNCard || c > MaxNCard {
		return CardPointNone
	}
	var rank = 1 + (int32(c)+1)%EachSuitN
	return CardPoint(rank)
}

func (c Card) RankString() string {
	var pt = c.Rank()
	if pt >= CardPoint3 && pt <= CardPoint10 {
		return strconv.Itoa(int(pt) + 2)
	}
	switch pt {
	case CardPointJ:
		return NameRankJ
	case CardPointQ:
		return NameRankQ
	case CardPointK:
		return NameRankK
	case CardPointA:
		return NameRankA
	case CardPoint2:
		return NameRank2
	}
	return "?"
}

func (c Card) IsJoker() bool {
	return c == CardBlackJoker || c == CardRedJoker
}

func (c Card) IsFaceCard() bool {
	var r = c.Rank()
	return r >= CardPointJ && r <= CardPointK
}

func (c Card) String() string {
	if c.IsJoker() {
		return c.SuitString()
	}
	return c.SuitString() + c.RankString()
}
