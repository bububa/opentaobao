package wdk

import (
	"sync"
)

// CancelRequest 结构体
type CancelRequest struct {
	// 子订单号列表
	SubBizOrderCodes []string `json:"sub_biz_order_codes,omitempty" xml:"sub_biz_order_codes>string,omitempty"`
	// 订单号
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 渠道来源
	SourceFrom string `json:"source_from,omitempty" xml:"source_from,omitempty"`
	// 出库单单据类型
	OutboundOrderType int64 `json:"outbound_order_type,omitempty" xml:"outbound_order_type,omitempty"`
}

var poolCancelRequest = sync.Pool{
	New: func() any {
		return new(CancelRequest)
	},
}

// GetCancelRequest() 从对象池中获取CancelRequest
func GetCancelRequest() *CancelRequest {
	return poolCancelRequest.Get().(*CancelRequest)
}

// ReleaseCancelRequest 释放CancelRequest
func ReleaseCancelRequest(v *CancelRequest) {
	v.SubBizOrderCodes = v.SubBizOrderCodes[:0]
	v.SourceOrderCode = ""
	v.WarehouseCode = ""
	v.SourceFrom = ""
	v.OutboundOrderType = 0
	poolCancelRequest.Put(v)
}
