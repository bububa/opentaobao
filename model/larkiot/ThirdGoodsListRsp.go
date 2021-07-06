package larkiot

// ThirdGoodsListRsp 结构体
type ThirdGoodsListRsp struct {
	// 卖品列表
	GoodsList []ThirdGoodsRsp `json:"goods_list,omitempty" xml:"goods_list>third_goods_rsp,omitempty"`
	// 影院内码
	CinemaLinkId string `json:"cinema_link_id,omitempty" xml:"cinema_link_id,omitempty"`
	// 最大数量
	MaxBuyCount int64 `json:"max_buy_count,omitempty" xml:"max_buy_count,omitempty"`
}
