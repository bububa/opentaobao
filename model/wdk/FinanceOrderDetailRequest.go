package wdk

import (
	"sync"
)

// FinanceOrderDetailRequest 结构体
type FinanceOrderDetailRequest struct {
	// 门店编码list
	ShopCodes []string `json:"shop_codes,omitempty" xml:"shop_codes>string,omitempty"`
	// 销售渠道
	SaleChannel string `json:"sale_channel,omitempty" xml:"sale_channel,omitempty"`
	// 销售来源
	SaleSource string `json:"sale_source,omitempty" xml:"sale_source,omitempty"`
	// 交易类型
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolFinanceOrderDetailRequest = sync.Pool{
	New: func() any {
		return new(FinanceOrderDetailRequest)
	},
}

// GetFinanceOrderDetailRequest() 从对象池中获取FinanceOrderDetailRequest
func GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
	return poolFinanceOrderDetailRequest.Get().(*FinanceOrderDetailRequest)
}

// ReleaseFinanceOrderDetailRequest 释放FinanceOrderDetailRequest
func ReleaseFinanceOrderDetailRequest(v *FinanceOrderDetailRequest) {
	v.ShopCodes = v.ShopCodes[:0]
	v.SaleChannel = ""
	v.SaleSource = ""
	v.TradeType = ""
	v.EndTime = ""
	v.StartTime = ""
	v.CurrentPage = 0
	v.PageSize = 0
	poolFinanceOrderDetailRequest.Put(v)
}
