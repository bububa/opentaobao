package alsc

import (
	"sync"
)

// AttachInfo 结构体
type AttachInfo struct {
	// 附加商品id
	OutAttachItemId string `json:"out_attach_item_id,omitempty" xml:"out_attach_item_id,omitempty"`
	// 附加商品名称
	OutAttachItemName string `json:"out_attach_item_name,omitempty" xml:"out_attach_item_name,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 附加商品实收金额
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 附加商品数量
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
	// 单价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 附加商品总金额
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
}

var poolAttachInfo = sync.Pool{
	New: func() any {
		return new(AttachInfo)
	},
}

// GetAttachInfo() 从对象池中获取AttachInfo
func GetAttachInfo() *AttachInfo {
	return poolAttachInfo.Get().(*AttachInfo)
}

// ReleaseAttachInfo 释放AttachInfo
func ReleaseAttachInfo(v *AttachInfo) {
	v.OutAttachItemId = ""
	v.OutAttachItemName = ""
	v.Unit = ""
	v.Weight = ""
	v.ActualFee = 0
	v.ItemCount = 0
	v.Price = 0
	v.TotalFee = 0
	poolAttachInfo.Put(v)
}
