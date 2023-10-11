package jym

// JymRecommendGoodsInfoDto 结构体
type JymRecommendGoodsInfoDto struct {
	// 推荐商品
	GoodsList []JymSingleGoodsDto `json:"goods_list,omitempty" xml:"goods_list>jym_single_goods_dto,omitempty"`
	// 交易猫游戏ID
	JymGameId string `json:"jym_game_id,omitempty" xml:"jym_game_id,omitempty"`
}
