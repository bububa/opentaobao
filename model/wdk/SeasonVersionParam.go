package wdk

// SeasonVersionParam 结构体
type SeasonVersionParam struct {
	// 参与换挡的门店列表
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 商家档期号
	OutSeasonId string `json:"out_season_id,omitempty" xml:"out_season_id,omitempty"`
}
