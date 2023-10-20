package wdk

import (
	"sync"
)

// ExtMap 结构体
type ExtMap struct {
	// 订单小号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 最晚拣货完成时间
	LatestPrepareTime string `json:"latest_prepare_time,omitempty" xml:"latest_prepare_time,omitempty"`
}

var poolExtMap = sync.Pool{
	New: func() any {
		return new(ExtMap)
	},
}

// GetExtMap() 从对象池中获取ExtMap
func GetExtMap() *ExtMap {
	return poolExtMap.Get().(*ExtMap)
}

// ReleaseExtMap 释放ExtMap
func ReleaseExtMap(v *ExtMap) {
	v.OrderNo = ""
	v.LatestPrepareTime = ""
	poolExtMap.Put(v)
}
