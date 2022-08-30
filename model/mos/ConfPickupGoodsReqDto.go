package mos

// ConfPickupGoodsReqDto 结构体
type ConfPickupGoodsReqDto struct {
	// 操作员
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 核销码
	ShortCode string `json:"short_code,omitempty" xml:"short_code,omitempty"`
	// OC交易号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 门店号
	OutBelongId string `json:"out_belong_id,omitempty" xml:"out_belong_id,omitempty"`
	// 专柜号
	OutStoreId string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`
}
