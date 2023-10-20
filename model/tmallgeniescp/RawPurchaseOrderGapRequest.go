package tmallgeniescp

import (
	"sync"
)

// RawPurchaseOrderGapRequest 结构体
type RawPurchaseOrderGapRequest struct {
	// 请求对象列表
	RawPurchaseOrderGapDTOs []RawPurchaseOrderGapDto `json:"raw_purchase_order_gap_d_t_os,omitempty" xml:"raw_purchase_order_gap_d_t_os>raw_purchase_order_gap_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolRawPurchaseOrderGapRequest = sync.Pool{
	New: func() any {
		return new(RawPurchaseOrderGapRequest)
	},
}

// GetRawPurchaseOrderGapRequest() 从对象池中获取RawPurchaseOrderGapRequest
func GetRawPurchaseOrderGapRequest() *RawPurchaseOrderGapRequest {
	return poolRawPurchaseOrderGapRequest.Get().(*RawPurchaseOrderGapRequest)
}

// ReleaseRawPurchaseOrderGapRequest 释放RawPurchaseOrderGapRequest
func ReleaseRawPurchaseOrderGapRequest(v *RawPurchaseOrderGapRequest) {
	v.RawPurchaseOrderGapDTOs = v.RawPurchaseOrderGapDTOs[:0]
	v.RequestExtendJson = ""
	poolRawPurchaseOrderGapRequest.Put(v)
}
