package inventory

import (
	"sync"
)

// PromiseList 结构体
type PromiseList struct {
	// 服务时效信息，发货时间比如2020-01-26；相对时间比如3，代表付款后3天内发货
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 履约仓code
	PerformStore string `json:"perform_store,omitempty" xml:"perform_store,omitempty"`
	// 发货时间类型，0-绝对时间，1-相对时间
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}

var poolPromiseList = sync.Pool{
	New: func() any {
		return new(PromiseList)
	},
}

// GetPromiseList() 从对象池中获取PromiseList
func GetPromiseList() *PromiseList {
	return poolPromiseList.Get().(*PromiseList)
}

// ReleasePromiseList 释放PromiseList
func ReleasePromiseList(v *PromiseList) {
	v.DeliveryTime = ""
	v.PerformStore = ""
	v.DeliveryType = 0
	poolPromiseList.Put(v)
}
