package wdk

import (
	"sync"
)

// AfterRefundOrderRequest 结构体
type AfterRefundOrderRequest struct {
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 业务子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退款信息
	AfterRefundOrderInfo *AfterRefundOrderInfo `json:"after_refund_order_info,omitempty" xml:"after_refund_order_info,omitempty"`
}

var poolAfterRefundOrderRequest = sync.Pool{
	New: func() any {
		return new(AfterRefundOrderRequest)
	},
}

// GetAfterRefundOrderRequest() 从对象池中获取AfterRefundOrderRequest
func GetAfterRefundOrderRequest() *AfterRefundOrderRequest {
	return poolAfterRefundOrderRequest.Get().(*AfterRefundOrderRequest)
}

// ReleaseAfterRefundOrderRequest 释放AfterRefundOrderRequest
func ReleaseAfterRefundOrderRequest(v *AfterRefundOrderRequest) {
	v.StoreId = ""
	v.SubBizOrderId = ""
	v.AfterRefundOrderInfo = nil
	poolAfterRefundOrderRequest.Put(v)
}
