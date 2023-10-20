package ascp

import (
	"sync"
)

// CancelConsignOrderRequest 结构体
type CancelConsignOrderRequest struct {
	// 交易子单号
	SubTradeIds []string `json:"sub_trade_ids,omitempty" xml:"sub_trade_ids>string,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 翱象的发货单号（翱象发货单号、交易主单号两个字段必须填一个）
	ConsignOrderCode string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty"`
	// 交易主单号（翱象发货单号、交易主单号两个字段必须填一个）
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 取消原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 业务请求Id，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolCancelConsignOrderRequest = sync.Pool{
	New: func() any {
		return new(CancelConsignOrderRequest)
	},
}

// GetCancelConsignOrderRequest() 从对象池中获取CancelConsignOrderRequest
func GetCancelConsignOrderRequest() *CancelConsignOrderRequest {
	return poolCancelConsignOrderRequest.Get().(*CancelConsignOrderRequest)
}

// ReleaseCancelConsignOrderRequest 释放CancelConsignOrderRequest
func ReleaseCancelConsignOrderRequest(v *CancelConsignOrderRequest) {
	v.SubTradeIds = v.SubTradeIds[:0]
	v.WarehouseCode = ""
	v.OwnerCode = ""
	v.ConsignOrderCode = ""
	v.TradeId = ""
	v.CancelReason = ""
	v.ExtendProps = ""
	v.RequestId = ""
	v.RequestTime = 0
	poolCancelConsignOrderRequest.Put(v)
}
