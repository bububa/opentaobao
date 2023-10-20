package ascp

import (
	"sync"
)

// OrderRdcCreate 结构体
type OrderRdcCreate struct {
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

var poolOrderRdcCreate = sync.Pool{
	New: func() any {
		return new(OrderRdcCreate)
	},
}

// GetOrderRdcCreate() 从对象池中获取OrderRdcCreate
func GetOrderRdcCreate() *OrderRdcCreate {
	return poolOrderRdcCreate.Get().(*OrderRdcCreate)
}

// ReleaseOrderRdcCreate 释放OrderRdcCreate
func ReleaseOrderRdcCreate(v *OrderRdcCreate) {
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolOrderRdcCreate.Put(v)
}
