package ascp

import (
	"sync"
)

// OrderSenderCreate 结构体
type OrderSenderCreate struct {
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人手机
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

var poolOrderSenderCreate = sync.Pool{
	New: func() any {
		return new(OrderSenderCreate)
	},
}

// GetOrderSenderCreate() 从对象池中获取OrderSenderCreate
func GetOrderSenderCreate() *OrderSenderCreate {
	return poolOrderSenderCreate.Get().(*OrderSenderCreate)
}

// ReleaseOrderSenderCreate 释放OrderSenderCreate
func ReleaseOrderSenderCreate(v *OrderSenderCreate) {
	v.Name = ""
	v.Mobile = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolOrderSenderCreate.Put(v)
}
