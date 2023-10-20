package trade

import (
	"sync"
)

// OrderDelivery 结构体
type OrderDelivery struct {
	// 配送开始时间
	DeliveryStartTime string `json:"delivery_start_time,omitempty" xml:"delivery_start_time,omitempty"`
	// 收货人姓名
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 配送结束时间
	DeliveryEndTime string `json:"delivery_end_time,omitempty" xml:"delivery_end_time,omitempty"`
	// 配送坐标
	DeliveryGeo string `json:"delivery_geo,omitempty" xml:"delivery_geo,omitempty"`
	// 配送地址
	DeliveryAddress string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	// 收货人电话
	ConsigneePhone string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
	// 配送员编码
	DelivererCode string `json:"deliverer_code,omitempty" xml:"deliverer_code,omitempty"`
	// 配送员电话
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 配送员姓名
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 配送费用金额
	DeliveryFee int64 `json:"delivery_fee,omitempty" xml:"delivery_fee,omitempty"`
	// 是否隐私号 0：非隐私号  1：隐私号
	PrivacyPhoneFlag int64 `json:"privacy_phone_flag,omitempty" xml:"privacy_phone_flag,omitempty"`
}

var poolOrderDelivery = sync.Pool{
	New: func() any {
		return new(OrderDelivery)
	},
}

// GetOrderDelivery() 从对象池中获取OrderDelivery
func GetOrderDelivery() *OrderDelivery {
	return poolOrderDelivery.Get().(*OrderDelivery)
}

// ReleaseOrderDelivery 释放OrderDelivery
func ReleaseOrderDelivery(v *OrderDelivery) {
	v.DeliveryStartTime = ""
	v.ConsigneeName = ""
	v.DeliveryEndTime = ""
	v.DeliveryGeo = ""
	v.DeliveryAddress = ""
	v.ConsigneePhone = ""
	v.DelivererCode = ""
	v.DelivererPhone = ""
	v.DelivererName = ""
	v.DeliveryFee = 0
	v.PrivacyPhoneFlag = 0
	poolOrderDelivery.Put(v)
}
