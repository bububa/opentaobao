package ascp

import (
	"sync"
)

// OrderReceiverEdit 结构体
type OrderReceiverEdit struct {
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人电话
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

var poolOrderReceiverEdit = sync.Pool{
	New: func() any {
		return new(OrderReceiverEdit)
	},
}

// GetOrderReceiverEdit() 从对象池中获取OrderReceiverEdit
func GetOrderReceiverEdit() *OrderReceiverEdit {
	return poolOrderReceiverEdit.Get().(*OrderReceiverEdit)
}

// ReleaseOrderReceiverEdit 释放OrderReceiverEdit
func ReleaseOrderReceiverEdit(v *OrderReceiverEdit) {
	v.Name = ""
	v.Mobile = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolOrderReceiverEdit.Put(v)
}
