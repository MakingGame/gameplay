// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHandCase struct {
	input    Cards
	expected CardPoint
}

func testCheckCardCategory(t *testing.T, tests []testHandCase, category HandsCategory) {
	for _, tc := range tests {
		var cards = tc.input
		sort.Sort(cards)
		//t.Logf("check cards %v %v", cards, category)
		var pt = cards.GetCategoryPoint(category)
		assert.Equal(t, pt, tc.expected)
	}
}

func TestCards_CardsFromStr(t *testing.T) {
	var ss = CardsFromStr("方3", "梅3")
	t.Logf("%v", ss)
}

func TestCards_Pair(t *testing.T) {
	tests := []testHandCase{
		{CardsFromStr("方3", "梅3"), CardPoint3},
		{CardsFromStr("方A", "红A"), CardPointA},
		{CardsFromStr("红A", "梅3"), 0},
	}
	testCheckCardCategory(t, tests, CategoryAA)

	tests2 := []testHandCase{
		{CardsFromStr("大王", "小王"), CardPointRJK},
	}
	testCheckCardCategory(t, tests2, CategoryRR)
}

func TestCards_Triplet(t *testing.T) {
	tests1 := []testHandCase{
		{CardsFromStr("方3", "梅3", "红3"), CardPoint3},
		{CardsFromStr("方A", "梅A", "红A"), CardPointA},
		{CardsFromStr("方3", "梅3", "红4"), 0},
		{CardsFromStr("方3", "红4"), 0},
	}
	testCheckCardCategory(t, tests1, CategoryAAA)

	tests2 := []testHandCase{
		{CardsFromStr("方4", "梅4", "红4", "红3"), CardPoint4},
		{CardsFromStr("方A", "梅A", "红A", "大王"), CardPointA}, //
		{CardsFromStr("方A", "梅A", "红A", "红A"), 0},
		{CardsFromStr("方A", "梅A", "红A"), 0},
	}
	testCheckCardCategory(t, tests2, CategoryAAAX)

	tests3 := []testHandCase{
		{CardsFromStr("方4", "梅4", "红4", "红3", "梅3"), CardPoint4},
		{CardsFromStr("方A", "梅A", "红A", "红4", "梅4"), CardPointA},
		{CardsFromStr("方A", "梅A", "红A", "红4", "梅3"), 0},
		{CardsFromStr("方A", "梅A", "红A", "红A", "梅3"), 0},
		{CardsFromStr("方A", "梅A", "红A", "红4"), 0},
	}
	testCheckCardCategory(t, tests3, CategoryAAAXX)
}

func TestCards_Sequence(t *testing.T) {
	tests1 := []testHandCase{
		{CardsFromStr("方3", "梅4", "红5", "红6", "梅7", "黑8", "黑9", "黑10", "黑J", "黑Q", "黑K"), CardPointK},
		{CardsFromStr("方10", "梅J", "红Q", "红K", "梅A"), CardPointA},
		{CardsFromStr("方10", "梅J", "红Q", "红K", "梅A", "方2"), 0},
		{CardsFromStr("梅J", "红Q", "红K", "梅A"), 0},
	}
	testCheckCardCategory(t, tests1, CategoryABCDE)

	tests2 := []testHandCase{
		{CardsFromStr("方A", "梅A", "红K", "黑K", "红Q", "黑Q"), CardPointA},
		{CardsFromStr("方3", "梅3", "红4", "黑4", "红5", "黑5", "红6", "黑6"), CardPoint6},
		{CardsFromStr("方3", "梅3", "红4", "黑4", "红5", "黑6"), 0},
		{CardsFromStr("方3", "梅3", "红4", "黑4"), 0},
		{CardsFromStr("方3", "梅3", "红3", "红3", "红4", "黑4"), 0},
	}
	testCheckCardCategory(t, tests2, CategoryAABBCC)
}

func TestCard_TripletsSequence(t *testing.T) {
	tests1 := []testHandCase{
		{CardsFromStr("方3", "梅3", "红3", "方4", "红4", "黑4", "方5", "红5", "黑5"), CardPoint5},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K", "方Q", "红Q", "黑Q", "方J", "红J", "黑J"), CardPointA},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K"), CardPointA},
		{CardsFromStr("方A", "梅A", "红A"), 0},
		{CardsFromStr("方3", "梅3", "红3", "方4", "红4"), 0},
		{CardsFromStr("方3", "梅3", "红3", "方4", "红4", "黑4", "方5", "红5", "黑6"), 0},
		{CardsFromStr("方2", "梅2", "红2", "方A", "梅A", "红A", "方K", "红K", "黑K"), 0},
	}
	testCheckCardCategory(t, tests1, CategoryAAABBB)

	tests2 := []testHandCase{
		{CardsFromStr("方3", "梅3", "红3", "方4", "红4", "黑4", "方5", "红5", "黑5", "黑6", "黑7", "黑8"), CardPoint5},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K", "方3", "红4"), CardPointA},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K", "方3", "红3"), 0},
		{CardsFromStr("方3", "梅3", "红3", "方5", "红5", "黑5", "方3", "红4"), 0},
	}
	testCheckCardCategory(t, tests2, CategoryAAABBBXY)

	tests3 := []testHandCase{
		{CardsFromStr("方3", "梅3", "红3", "方4", "红4", "黑4", "方5", "红5", "黑5", "黑6", "黑6", "黑7", "红7", "黑8", "红8"), CardPoint5},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K", "方3", "黑3", "红4", "黑4"), CardPointA},
		{CardsFromStr("方A", "梅A", "红A", "方K", "红K", "黑K", "方3", "黑4", "红4", "黑6"), 0},
		{CardsFromStr("方A", "梅A", "红A", "方3", "红3"), 0},
		{CardsFromStr("方3", "梅3", "红3", "方5", "红5", "黑5", "方3", "红4"), 0},
	}
	testCheckCardCategory(t, tests3, CategoryAAABBBXXYY)
}

func TestCard_Quadplex(t *testing.T) {
	tests1 := []testHandCase{
		{CardsFromStr("方A", "梅A", "红A", "黑A", "黑K", "黑Q"), CardPointA},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4", "黑5"), CardPoint3},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "方4", "梅4", "红4", "黑4", "黑2", "红5", "黑6", "红7"), 0},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4", "黑4"), 0},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4"), 0},
		{CardsFromStr("方3", "梅3", "红3", "黑3"), 0},
	}
	testCheckCardCategory(t, tests1, CategoryAAAAXY)

	tests2 := []testHandCase{
		{CardsFromStr("方A", "梅A", "红A", "黑A", "黑K", "红K", "黑Q", "红Q"), CardPointA},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4", "红4", "黑5", "红5"), CardPoint3},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4", "红4", "黑5", "红6"), 0},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "黑4", "红5", "黑6", "红7"), 0},
		{CardsFromStr("方3", "梅3", "红3", "黑3", "方4", "梅4", "红4", "黑4"), 0},
	}
	testCheckCardCategory(t, tests2, CategoryAAAAXXYY)
}
