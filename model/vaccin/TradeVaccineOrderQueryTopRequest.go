package vaccin

import (
	"sync"
)

// TradeVaccineOrderQueryTopRequest 结构体
type TradeVaccineOrderQueryTopRequest struct {
	// 卖家ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 业务订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolTradeVaccineOrderQueryTopRequest = sync.Pool{
	New: func() any {
		return new(TradeVaccineOrderQueryTopRequest)
	},
}

// GetTradeVaccineOrderQueryTopRequest() 从对象池中获取TradeVaccineOrderQueryTopRequest
func GetTradeVaccineOrderQueryTopRequest() *TradeVaccineOrderQueryTopRequest {
	return poolTradeVaccineOrderQueryTopRequest.Get().(*TradeVaccineOrderQueryTopRequest)
}

// ReleaseTradeVaccineOrderQueryTopRequest 释放TradeVaccineOrderQueryTopRequest
func ReleaseTradeVaccineOrderQueryTopRequest(v *TradeVaccineOrderQueryTopRequest) {
	v.MerchantId = ""
	v.BizOrderId = 0
	poolTradeVaccineOrderQueryTopRequest.Put(v)
}
