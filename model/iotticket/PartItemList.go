package iotticket

import (
	"sync"
)

// PartItemList 结构体
type PartItemList struct {
	// 配件编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 支付角色：merchant-商家记账；customer-客户付费
	PayRole string `json:"pay_role,omitempty" xml:"pay_role,omitempty"`
}

var poolPartItemList = sync.Pool{
	New: func() any {
		return new(PartItemList)
	},
}

// GetPartItemList() 从对象池中获取PartItemList
func GetPartItemList() *PartItemList {
	return poolPartItemList.Get().(*PartItemList)
}

// ReleasePartItemList 释放PartItemList
func ReleasePartItemList(v *PartItemList) {
	v.ItemCode = ""
	v.PayRole = ""
	poolPartItemList.Put(v)
}
