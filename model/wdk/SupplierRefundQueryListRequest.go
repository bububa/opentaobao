package wdk

import (
	"sync"
)

// SupplierRefundQueryListRequest 结构体
type SupplierRefundQueryListRequest struct {
	// 盒马主订单id
	MainBizOrderIds []int64 `json:"main_biz_order_ids,omitempty" xml:"main_biz_order_ids>int64,omitempty"`
	// 盒马子订单id
	SubBizOrderIds []int64 `json:"sub_biz_order_ids,omitempty" xml:"sub_biz_order_ids>int64,omitempty"`
	// 退款单id
	RefundIds []int64 `json:"refund_ids,omitempty" xml:"refund_ids>int64,omitempty"`
	// 售卖商场code
	SourceMerchantCode string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
	// 渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolSupplierRefundQueryListRequest = sync.Pool{
	New: func() any {
		return new(SupplierRefundQueryListRequest)
	},
}

// GetSupplierRefundQueryListRequest() 从对象池中获取SupplierRefundQueryListRequest
func GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
	return poolSupplierRefundQueryListRequest.Get().(*SupplierRefundQueryListRequest)
}

// ReleaseSupplierRefundQueryListRequest 释放SupplierRefundQueryListRequest
func ReleaseSupplierRefundQueryListRequest(v *SupplierRefundQueryListRequest) {
	v.MainBizOrderIds = v.MainBizOrderIds[:0]
	v.SubBizOrderIds = v.SubBizOrderIds[:0]
	v.RefundIds = v.RefundIds[:0]
	v.SourceMerchantCode = ""
	v.OrderFrom = 0
	poolSupplierRefundQueryListRequest.Put(v)
}
