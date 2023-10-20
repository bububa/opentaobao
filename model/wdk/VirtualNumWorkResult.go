package wdk

import (
	"sync"
)

// VirtualNumWorkResult 结构体
type VirtualNumWorkResult struct {
	// 淘鲜达交易单号
	SourceOrderId string `json:"source_order_id,omitempty" xml:"source_order_id,omitempty"`
	// 查询类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 返回虚拟号
	VirtualNumber string `json:"virtual_number,omitempty" xml:"virtual_number,omitempty"`
	// 订阅id
	SubId string `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
}

var poolVirtualNumWorkResult = sync.Pool{
	New: func() any {
		return new(VirtualNumWorkResult)
	},
}

// GetVirtualNumWorkResult() 从对象池中获取VirtualNumWorkResult
func GetVirtualNumWorkResult() *VirtualNumWorkResult {
	return poolVirtualNumWorkResult.Get().(*VirtualNumWorkResult)
}

// ReleaseVirtualNumWorkResult 释放VirtualNumWorkResult
func ReleaseVirtualNumWorkResult(v *VirtualNumWorkResult) {
	v.SourceOrderId = ""
	v.Type = ""
	v.VirtualNumber = ""
	v.SubId = ""
	poolVirtualNumWorkResult.Put(v)
}
