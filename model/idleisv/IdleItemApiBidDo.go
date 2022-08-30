package idleisv

// IdleItemApiBidDo 结构体
type IdleItemApiBidDo struct {
	// 拍卖商品id
	BidItemId int64 `json:"bid_item_id,omitempty" xml:"bid_item_id,omitempty"`
	// 拍卖开始时间,传入豪秒时间戳(底层精度为分钟)
	BidStartTime int64 `json:"bid_start_time,omitempty" xml:"bid_start_time,omitempty"`
	// 拍卖结束时间,传入豪秒时间戳(底层精度为分钟)
	BidEndTime int64 `json:"bid_end_time,omitempty" xml:"bid_end_time,omitempty"`
	// 买家参拍保证金金额，单位：分
	BidBail int64 `json:"bid_bail,omitempty" xml:"bid_bail,omitempty"`
	// 出价幅度金额(起拍价使用reserve_price)，单位：分
	BidStep int64 `json:"bid_step,omitempty" xml:"bid_step,omitempty"`
	// 当前出价金额, 最高出价, 单位分
	CurrentBidPrice int64 `json:"current_bid_price,omitempty" xml:"current_bid_price,omitempty"`
	// 总出价次数
	BidCount int64 `json:"bid_count,omitempty" xml:"bid_count,omitempty"`
	// 最后两分钟被出价的延迟次数
	DelayCount int64 `json:"delay_count,omitempty" xml:"delay_count,omitempty"`
}
