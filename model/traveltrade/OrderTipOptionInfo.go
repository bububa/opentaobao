package traveltrade

import (
	"sync"
)

// OrderTipOptionInfo 结构体
type OrderTipOptionInfo struct {
	// select对应的text值
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// select对应的value值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolOrderTipOptionInfo = sync.Pool{
	New: func() any {
		return new(OrderTipOptionInfo)
	},
}

// GetOrderTipOptionInfo() 从对象池中获取OrderTipOptionInfo
func GetOrderTipOptionInfo() *OrderTipOptionInfo {
	return poolOrderTipOptionInfo.Get().(*OrderTipOptionInfo)
}

// ReleaseOrderTipOptionInfo 释放OrderTipOptionInfo
func ReleaseOrderTipOptionInfo(v *OrderTipOptionInfo) {
	v.Text = ""
	v.Value = ""
	poolOrderTipOptionInfo.Put(v)
}
