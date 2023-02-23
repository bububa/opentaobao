package caipiao

// LotteryWangcaiSellerGoodsInfo 结构体
type LotteryWangcaiSellerGoodsInfo struct {
	// 活动结束时间
	ActEndTime string `json:"act_end_time,omitempty" xml:"act_end_time,omitempty"`
	// 活动开始时间
	ActBeginTime string `json:"act_begin_time,omitempty" xml:"act_begin_time,omitempty"`
	// 商品id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 赠送类型：0：满就送，1：好评送，2：分享送，3：游戏送，4：收藏送
	PresentType int64 `json:"present_type,omitempty" xml:"present_type,omitempty"`
	// 彩种id
	LotteryTypeId int64 `json:"lottery_type_id,omitempty" xml:"lottery_type_id,omitempty"`
}
