package qimen

import (
	"sync"
)

// WarehouseInfo 结构体
type WarehouseInfo struct {
	// 奇门仓储字段
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 奇门仓储字段
	WarehouseName string `json:"warehouseName,omitempty" xml:"warehouseName,omitempty"`
	// 奇门仓储字段
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 奇门仓储字段
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 奇门仓储字段
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 奇门仓储字段
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 奇门仓储字段
	DetailAddress string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// 奇门仓储字段
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奇门仓储字段
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 奇门仓储字段
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 奇门仓储字段
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolWarehouseInfo = sync.Pool{
	New: func() any {
		return new(WarehouseInfo)
	},
}

// GetWarehouseInfo() 从对象池中获取WarehouseInfo
func GetWarehouseInfo() *WarehouseInfo {
	return poolWarehouseInfo.Get().(*WarehouseInfo)
}

// ReleaseWarehouseInfo 释放WarehouseInfo
func ReleaseWarehouseInfo(v *WarehouseInfo) {
	v.WarehouseCode = ""
	v.WarehouseName = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	v.Name = ""
	v.Tel = ""
	v.Mobile = ""
	v.Status = ""
	poolWarehouseInfo.Put(v)
}
