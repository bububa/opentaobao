package perfect

import (
	"sync"
)

// OutboundCancelRequest 结构体
type OutboundCancelRequest struct {
	// 出库子单
	SubOrderCodes []string `json:"sub_order_codes,omitempty" xml:"sub_order_codes>string,omitempty"`
	// 出库单
	OutboundOrderCode string `json:"outbound_order_code,omitempty" xml:"outbound_order_code,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolOutboundCancelRequest = sync.Pool{
	New: func() any {
		return new(OutboundCancelRequest)
	},
}

// GetOutboundCancelRequest() 从对象池中获取OutboundCancelRequest
func GetOutboundCancelRequest() *OutboundCancelRequest {
	return poolOutboundCancelRequest.Get().(*OutboundCancelRequest)
}

// ReleaseOutboundCancelRequest 释放OutboundCancelRequest
func ReleaseOutboundCancelRequest(v *OutboundCancelRequest) {
	v.SubOrderCodes = v.SubOrderCodes[:0]
	v.OutboundOrderCode = ""
	v.WarehouseCode = ""
	poolOutboundCancelRequest.Put(v)
}
