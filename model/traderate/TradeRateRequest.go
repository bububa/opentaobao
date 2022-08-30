package traderate

// TradeRateRequest 结构体
type TradeRateRequest struct {
	// 评价创建时间,格式:yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 子订单ID
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
}
