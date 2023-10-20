package ascp

import (
	"sync"
)

// OrderSenderEdit 结构体
type OrderSenderEdit struct {
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

var poolOrderSenderEdit = sync.Pool{
	New: func() any {
		return new(OrderSenderEdit)
	},
}

// GetOrderSenderEdit() 从对象池中获取OrderSenderEdit
func GetOrderSenderEdit() *OrderSenderEdit {
	return poolOrderSenderEdit.Get().(*OrderSenderEdit)
}

// ReleaseOrderSenderEdit 释放OrderSenderEdit
func ReleaseOrderSenderEdit(v *OrderSenderEdit) {
	v.Name = ""
	v.Mobile = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolOrderSenderEdit.Put(v)
}
