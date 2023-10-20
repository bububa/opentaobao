package waybill

import (
	"sync"
)

// WaybillApplyNewRequest 结构体
type WaybillApplyNewRequest struct {
	// 订单数据
	TradeOrderInfoCols []TradeOrderInfo `json:"trade_order_info_cols,omitempty" xml:"trade_order_info_cols>trade_order_info,omitempty"`
	// 物流服务商编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 收\发货地址
	ShippingAddress *WaybillAddress `json:"shipping_address,omitempty" xml:"shipping_address,omitempty"`
}

var poolWaybillApplyNewRequest = sync.Pool{
	New: func() any {
		return new(WaybillApplyNewRequest)
	},
}

// GetWaybillApplyNewRequest() 从对象池中获取WaybillApplyNewRequest
func GetWaybillApplyNewRequest() *WaybillApplyNewRequest {
	return poolWaybillApplyNewRequest.Get().(*WaybillApplyNewRequest)
}

// ReleaseWaybillApplyNewRequest 释放WaybillApplyNewRequest
func ReleaseWaybillApplyNewRequest(v *WaybillApplyNewRequest) {
	v.TradeOrderInfoCols = v.TradeOrderInfoCols[:0]
	v.CpCode = ""
	v.ShippingAddress = nil
	poolWaybillApplyNewRequest.Put(v)
}
