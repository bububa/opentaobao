package jym

import (
	"sync"
)

// DeliveryNotifyDto 结构体
type DeliveryNotifyDto struct {
	// 订单ID
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 发货信息通知
	DeliveryMsg string `json:"delivery_msg,omitempty" xml:"delivery_msg,omitempty"`
	// 发货失败原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	// 发货类型：1：发货中， 2：发货成功 3.发货失败 4.发货通知
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}

var poolDeliveryNotifyDto = sync.Pool{
	New: func() any {
		return new(DeliveryNotifyDto)
	},
}

// GetDeliveryNotifyDto() 从对象池中获取DeliveryNotifyDto
func GetDeliveryNotifyDto() *DeliveryNotifyDto {
	return poolDeliveryNotifyDto.Get().(*DeliveryNotifyDto)
}

// ReleaseDeliveryNotifyDto 释放DeliveryNotifyDto
func ReleaseDeliveryNotifyDto(v *DeliveryNotifyDto) {
	v.BizOrderId = ""
	v.DeliveryMsg = ""
	v.ReasonId = 0
	v.DeliveryType = 0
	poolDeliveryNotifyDto.Put(v)
}
