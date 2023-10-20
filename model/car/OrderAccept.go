package car

import (
	"sync"
)

// OrderAccept 结构体
type OrderAccept struct {
	// 拒单原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 阿里旅行用车订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 服务商ID,阿里提供
	ProviderId string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 服务商订单ID,没有可不传
	ThirdOrderId string `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
	// 可选，卖家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 接单司机纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 接单司机经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 0确认接单 1无法接单
	ConfirmType int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 0:接送机 1：实时打车 2：租车(不传值默认为0)
	UseType int64 `json:"use_type,omitempty" xml:"use_type,omitempty"`
	// 接单时间毫秒数
	AcceptTime int64 `json:"accept_time,omitempty" xml:"accept_time,omitempty"`
}

var poolOrderAccept = sync.Pool{
	New: func() any {
		return new(OrderAccept)
	},
}

// GetOrderAccept() 从对象池中获取OrderAccept
func GetOrderAccept() *OrderAccept {
	return poolOrderAccept.Get().(*OrderAccept)
}

// ReleaseOrderAccept 释放OrderAccept
func ReleaseOrderAccept(v *OrderAccept) {
	v.Message = ""
	v.OrderId = ""
	v.ProviderId = ""
	v.ThirdOrderId = ""
	v.SellerId = ""
	v.Latitude = ""
	v.Longitude = ""
	v.ConfirmType = 0
	v.UseType = 0
	v.AcceptTime = 0
	poolOrderAccept.Put(v)
}
