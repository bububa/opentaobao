package wlbimports

import (
	"sync"
)

// LocOrder 结构体
type LocOrder struct {
	// 物流承运商
	Carrier string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 物流订单状态编码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 物流订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 物流运单号
	TrackingNo string `json:"tracking_no,omitempty" xml:"tracking_no,omitempty"`
	// 重量单位
	WeightUnit string `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	// 费用币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 订单状态中文描述
	StatusCodeDesc string `json:"status_code_desc,omitempty" xml:"status_code_desc,omitempty"`
	// 运费
	ShippingFee int64 `json:"shipping_fee,omitempty" xml:"shipping_fee,omitempty"`
	// 交易订单号
	TradeId int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 关税
	CustomsFee int64 `json:"customs_fee,omitempty" xml:"customs_fee,omitempty"`
	// 重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolLocOrder = sync.Pool{
	New: func() any {
		return new(LocOrder)
	},
}

// GetLocOrder() 从对象池中获取LocOrder
func GetLocOrder() *LocOrder {
	return poolLocOrder.Get().(*LocOrder)
}

// ReleaseLocOrder 释放LocOrder
func ReleaseLocOrder(v *LocOrder) {
	v.Carrier = ""
	v.StatusCode = ""
	v.OrderCode = ""
	v.TrackingNo = ""
	v.WeightUnit = ""
	v.Currency = ""
	v.StatusCodeDesc = ""
	v.ShippingFee = 0
	v.TradeId = 0
	v.CustomsFee = 0
	v.Weight = 0
	poolLocOrder.Put(v)
}
