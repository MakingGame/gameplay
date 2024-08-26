// Copyright © 2021 ichenq@gmail.com All rights reserved.
// See accompanying files LICENSE.txt

package ddz

const (
	MinNCard  = 1  // 最小牌面值
	MaxNCard  = 54 // 最大牌面值
	EachSuitN = 13 // 每花色牌数
)

const (
	NameBlackJoker = "小王"
	NameRedJoker   = "大王"
	NameDiamond    = "方"
	NameClub       = "梅"
	NameHeart      = "红"
	NameSpade      = "黑"
	NameRankJ      = "J"
	NameRankQ      = "Q"
	NameRankK      = "K"
	NameRankA      = "A"
	NameRank2      = "2"
)

// 牌的点数
type CardPoint int32

const (
	CardPointNone CardPoint = 0
	CardPoint3    CardPoint = 1
	CardPoint4    CardPoint = 2
	CardPoint5    CardPoint = 3
	CardPoint6    CardPoint = 4
	CardPoint7    CardPoint = 5
	CardPoint8    CardPoint = 6
	CardPoint9    CardPoint = 7
	CardPoint10   CardPoint = 8
	CardPointJ    CardPoint = 9
	CardPointQ    CardPoint = 10
	CardPointK    CardPoint = 11
	CardPointA    CardPoint = 12
	CardPoint2    CardPoint = 13
	CardPointBJK  CardPoint = 14
	CardPointRJK  CardPoint = 15
)

// 花色
type CardSuit int32

const (
	SuitNone           CardSuit = 0
	CardSuitDiamonds   CardSuit = 1
	CardSuitClubs      CardSuit = 2
	CardSuitHearts     CardSuit = 3
	CardSuitSpades     CardSuit = 4
	CardSuitBlackJoker CardSuit = 5
	CardSuitRedJoker   CardSuit = 6
)

const (
	// 方块3 - 10
	CardDiamond3  = 1
	CardDiamond4  = 2
	CardDiamond5  = 3
	CardDiamond6  = 4
	CardDiamond7  = 5
	CardDiamond8  = 6
	CardDiamond9  = 7
	CardDiamond10 = 8
	CardDiamondJ  = 9
	CardDiamondQ  = 10
	CardDiamondK  = 11
	CardDiamondA  = 12
	CardDiamond2  = 13

	// 梅花3 - 10
	CardClub3  = 14
	CardClub4  = 15
	CardClub5  = 16
	CardClub6  = 17
	CardClub7  = 18
	CardClub8  = 19
	CardClub9  = 20
	CardClub10 = 21
	CardClubJ  = 22
	CardClubQ  = 23
	CardClubK  = 24
	CardClubA  = 25
	CardClub2  = 26

	// 红心3 - 10
	CardHeart3  = 27
	CardHeart4  = 28
	CardHeart5  = 29
	CardHeart6  = 30
	CardHeart7  = 31
	CardHeart8  = 32
	CardHeart9  = 33
	CardHeart10 = 34
	CardHeartJ  = 35
	CardHeartQ  = 36
	CardHeartK  = 37
	CardHeartA  = 38
	CardHeart2  = 39

	// 黑桃3 - 10
	CardSpade3  = 40
	CardSpade4  = 41
	CardSpade5  = 42
	CardSpade6  = 43
	CardSpade7  = 44
	CardSpade8  = 45
	CardSpade9  = 46
	CardSpade10 = 47
	CardSpadeJ  = 48
	CardSpadeQ  = 49
	CardSpadeK  = 50
	CardSpadeA  = 51
	CardSpade2  = 52

	// 大小王
	CardBlackJoker = 53
	CardRedJoker   = 54
)

type HandsCategory int32

const (
	CategoryNone       HandsCategory = 0
	CategoryA          HandsCategory = 1  // 单张
	CategoryAA         HandsCategory = 2  // 对子
	CategoryAAA        HandsCategory = 3  // 三张
	CategoryAAAX       HandsCategory = 4  // 三带一
	CategoryAAAXX      HandsCategory = 5  // 三带对
	CategoryABCDE      HandsCategory = 6  // 顺子
	CategoryAABB       HandsCategory = 7  // 2连对
	CategoryAABBCC     HandsCategory = 8  // 3连对及以上
	CategoryAAABBB     HandsCategory = 9  // 3连牌
	CategoryAAABBBXY   HandsCategory = 10 // 3连牌带单牌
	CategoryAAABBBXXYY HandsCategory = 11 // 3连牌带对子
	CategoryAAAAXY     HandsCategory = 12 // 4带单牌
	CategoryAAAAXXYY   HandsCategory = 13 // 4带对子
	CategoryAAAA       HandsCategory = 14 // 炸弹
	CategoryRR         HandsCategory = 15 // 王炸
)
