package btrip

// OpenApiZzdSearchRq 结构体
type OpenApiZzdSearchRq struct {
	// 第三方企业ID
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 结算结束日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第三方用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 商旅申请单号
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 第几页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 第三方交易号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 结算开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}
