// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard_MakeCard(t *testing.T) {
	var tests = []struct {
		input1   CardSuit
		input2   CardPoint
		expected int32
	}{
		{CardSuitDiamonds, CardPoint3, CardDiamond3},
		{CardSuitClubs, CardPoint3, CardClub3},
		{CardSuitHearts, CardPoint2, CardHeart2},
		{CardSuitSpades, CardPointA, CardSpadeA},
		{CardSuitBlackJoker, 0, CardBlackJoker},
		{CardSuitRedJoker, 0, CardRedJoker},
	}
	for _, tc := range tests {
		var c = MakeCard(tc.input1, tc.input2)
		assert.Equal(t, int32(c), tc.expected)
	}
}

func TestCard_CardFromStr(t *testing.T) {
	var tests = []struct {
		input    string
		expected int32
	}{
		{NameRedJoker, CardRedJoker},
		{NameBlackJoker, CardBlackJoker},
		{NameDiamond + "3", CardDiamond3},
		{NameDiamond + "2", CardDiamond2},
		{NameDiamond + "A", CardDiamondA},
		{NameDiamond + "Q", CardDiamondQ},
		{NameSpade + "A", CardSpadeA},
		{NameSpade + "Q", CardSpadeQ},
	}
	for _, tc := range tests {
		var c = CardFromStr(tc.input)
		assert.Equal(t, int32(c), tc.expected)
	}
}
