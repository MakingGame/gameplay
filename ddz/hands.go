// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"sort"
)

// 打出的一副手牌
type Hands struct {
	Category HandsCategory
	Cards    []int32
}

// 一副手牌，已经按牌面排序
type Cards []Card

func CardsFromStr(ss ...string) Cards {
	if len(ss) == 0 {
		panic("no cards")
	}
	var cc = make(Cards, 0, len(ss))
	for _, s := range ss {
		var c = CardFromStr(s)
		if c == 0 {
			panic("invalid card " + s)
		}
		cc = append(cc, c)
	}
	return cc
}

// 可以组成连的牌[3 - A]
func IsSeqCardRank(r CardPoint) bool {
	return r >= CardPoint3 && r <= CardPointA
}

func (x Cards) Len() int           { return len(x) }
func (x Cards) Less(i, j int) bool { return x[i].Rank() > x[j].Rank() }
func (x Cards) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x Cards) Contains(c Card) bool {
	var i = sort.Search(len(x), func(i int) bool {
		return x[i] >= c
	})
	return i < len(x) && x[i] == c
}

// 单牌 CategoryA
func (x Cards) IsSingle() bool {
	return len(x) == 1
}

func (x Cards) SinglePoint() CardPoint {
	return x[0].Rank()
}

// 对子 CategoryAA
func (x Cards) IsPair() bool {
	if len(x) == 2 {
		return x[0].Rank() == x[1].Rank()
	}
	return false
}

func (x Cards) PairPoint() CardPoint {
	return x[0].Rank()
}

// 火箭 CategoryRR
func (x Cards) IsRocket() bool {
	// 已经按牌面值排序
	return len(x) == 2 &&
		x[0] == CardRedJoker &&
		x[1] == CardBlackJoker
}

// 炸弹 CategoryAAAA
func (x Cards) IsBomb() bool {
	if len(x) == 4 {
		var r = x[0].Rank()
		return r == x[1].Rank() && r == x[2].Rank() && r == x[3].Rank()
	}
	return false
}

func (x Cards) BombPoint() CardPoint {
	return x[0].Rank()
}

// 三张不带 CategoryAAA
func (x Cards) IsTripletRaw() bool {
	if len(x) == 3 {
		var r = x[0].Rank()
		return r == x[1].Rank() && r == x[2].Rank()
	}
	return false
}

// 三带一 CategoryAAAX
func (x Cards) IsTripletAttachSingle() bool {
	if len(x) == 4 {
		var r0 = x[0].Rank()
		var r1 = x[1].Rank()
		var r2 = x[2].Rank()
		var r3 = x[3].Rank()
		// 前3个相等或者后3个相等
		return (r0 == r1 && r1 == r2 && r0 != r3) ||
			(r0 != r1 && r1 == r2 && r2 == r3)
	}
	return false
}

// 三带对 CategoryAAAXX
func (x Cards) IsTripletAttachPair() bool {
	if len(x) == 5 {
		var r0 = x[0].Rank()
		var r1 = x[1].Rank()
		var r2 = x[2].Rank()
		var r3 = x[3].Rank()
		var r4 = x[4].Rank()
		// 前3个相等且后2个相等，或者前2个相等且后3个相等
		return (r0 == r1 && r1 == r2 && r2 != r3 && r3 == r4) ||
			(r0 == r1 && r1 != r2 && r2 == r3 && r3 == r4)
	}
	return false
}

func (x Cards) TripletPoint() CardPoint {
	var n = len(x)
	var r = x[0].Rank()
	if n == 3 {
		return r
	}
	// 前3个相等，点数就是x[0]的点数
	if r == x[1].Rank() && r == x[2].Rank() {
		return r
	}
	// 后3个相等，点数就是x[n-1]的点数
	return x[n-1].Rank()
}

// 顺子 CategoryABCDE
func (x Cards) IsSequence() bool {
	if len(x) < 5 {
		return false
	}
	var r = x[0].Rank() // 左边是最大的
	if !IsSeqCardRank(r) {
		return false
	}
	for i := 1; i < len(x); i++ {
		var v = x[i].Rank()
		if !(v > 0 && v+1 == r) {
			return false
		}
		r = v
	}
	return true
}

func (x Cards) SequencePoint() CardPoint {
	return x[0].Rank() // 左边是最大的
}

// 连对 CategoryAABBCC
func (x Cards) IsSequencePair(atLeast int) bool {
	if atLeast <= 0 {
		atLeast = 3
	}
	if n := len(x); n < atLeast*2 || n%2 != 0 {
		return false
	}
	var r = x[0].Rank()
	if !IsSeqCardRank(r) {
		return false // 判断最大的rank
	}
	if r != x[1].Rank() {
		return false
	}
	for i := 2; i < len(x); i += 2 {
		var a = x[i].Rank()
		var b = x[i+1].Rank()
		if !(a > 0 && a == b && a+1 == r) {
			return false
		}
		r = a
	}
	return true
}

// 飞机（不带） CategoryAAABBB
func (x Cards) IsSequenceTripletsRaw() bool {
	if len(x)%3 != 0 || len(x)/3 < 2 {
		return false
	}
	var r = x[0].Rank()
	if !IsSeqCardRank(r) {
		return false // 判断最大的rank
	}
	if r != x[1].Rank() || r != x[2].Rank() {
		return false
	}
	for i := 3; i < len(x); i += 3 {
		var a = x[i].Rank()
		var b = x[i+1].Rank()
		var c = x[i+2].Rank()
		if !(a > 0 && a == b && a == c && a+1 == r) {
			return false
		}
		r = a
	}
	return true
}

// 飞机（带单张）CategoryAAABBBXY
func (x Cards) IsSequenceTripletsAttachSingle() bool {
	if len(x)%4 != 0 || len(x)/4 < 2 {
		return false
	}
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}

	var nTriplet, nSingle int
	var lastTriplet CardPoint
	for i := 0; i < len(array); i++ {
		var r = CardPoint(i)
		switch array[i] {
		case 0:
		case 1:
			nSingle++
		case 3:
			nTriplet++
			if !IsSeqCardRank(r) {
				return false
			}
			if lastTriplet > 0 {
				if r != lastTriplet+1 {
					return false
				}
			}
			lastTriplet = r

		default:
			return false
		}
	}
	return nTriplet == nSingle
}

// 飞机（带对子） CategoryAAABBBXXYY
func (x Cards) IsSequenceTripletsAttachPair() bool {
	if len(x)%5 != 0 || len(x)/5 < 2 {
		return false
	}
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}
	var nTriplet, nPair int
	var lastTriplet CardPoint
	for i := 0; i < len(array); i++ {
		var r = CardPoint(i)
		switch array[i] {
		case 0:
		case 2:
			nPair++
		case 3:
			nTriplet++
			if !IsSeqCardRank(r) {
				return false
			}
			if lastTriplet > 0 {
				if r != lastTriplet+1 {
					return false
				}
			}
			lastTriplet = r
		default:
			return false
		}
	}
	return nTriplet == nPair
}

func (x Cards) SequenceTripletsPoint() CardPoint {
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}
	// 按牌面值由大到小遍历
	for i := len(array) - 1; i >= 0; i-- {
		switch array[i] {
		case 3:
			return CardPoint(i) // 第一个AAA值最大
		}
	}
	return 0
}

// 四带二 CategoryAAAAXY
func (x Cards) IsQuadplexAttachSingles() bool {
	if len(x) != 6 {
		return false
	}
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}
	var nQuad, nSingle int
	for i := 0; i < len(array); i++ {
		switch array[i] {
		case 0:
		case 1:
			nSingle++
		case 4:
			nQuad++
		default:
			return false
		}
	}
	return nQuad == 1 && nSingle == 2
}

// 四带两对 CategoryAAAAXXYY
func (x Cards) IsQuadplexAttachPairs() bool {
	if len(x) != 8 {
		return false
	}
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}
	var nQuad, nPair int
	for i := 0; i < len(array); i++ {
		switch array[i] {
		case 0:
		case 2:
			nPair++
		case 4:
			nQuad++
		default:
			return false
		}
	}
	return nQuad == 1 && nPair == 2
}

func (x Cards) QuadplexPoint() CardPoint {
	var array [CardPointRJK + 1]int32
	for i := 0; i < len(x); i++ {
		array[x[i].Rank()]++
	}
	for i := len(array) - 1; i >= 0; i-- {
		switch array[i] {
		case 4:
			return CardPoint(i)
		}
	}
	return 0
}

// 获取牌型的点数，没有成牌型则返回0
func (x Cards) GetCategoryPoint(category HandsCategory) CardPoint {
	switch category {
	case CategoryA:
		if x.IsSingle() {
			return x.SinglePoint()
		}
	case CategoryAA:
		if x.IsPair() {
			return x.PairPoint()
		}
	case CategoryAAA:
		if x.IsTripletRaw() {
			return x.TripletPoint()
		}
	case CategoryAAAA:
		if x.IsBomb() {
			return x.BombPoint()
		}
	case CategoryAAAX:
		if x.IsTripletAttachSingle() {
			return x.TripletPoint()
		}
	case CategoryAAAXX:
		if x.IsTripletAttachPair() {
			return x.TripletPoint()
		}
	case CategoryABCDE:
		if x.IsSequence() {
			return x.SequencePoint()
		}
	case CategoryAABB:
		if x.IsSequencePair(2) {
			return x.SequencePoint()
		}
	case CategoryAABBCC:
		if x.IsSequencePair(3) {
			return x.SequencePoint()
		}
	case CategoryAAABBB:
		if x.IsSequenceTripletsRaw() {
			return x.SequencePoint()
		}
	case CategoryAAABBBXY:
		if x.IsSequenceTripletsAttachSingle() {
			return x.SequenceTripletsPoint()
		}
	case CategoryAAABBBXXYY:
		if x.IsSequenceTripletsAttachPair() {
			return x.SequenceTripletsPoint()
		}
	case CategoryAAAAXY:
		if x.IsQuadplexAttachSingles() {
			return x.QuadplexPoint()
		}
	case CategoryAAAAXXYY:
		if x.IsQuadplexAttachPairs() {
			return x.QuadplexPoint()
		}
	case CategoryRR:
		if x.IsRocket() {
			return CardPointRJK
		}
	}
	return 0
}
