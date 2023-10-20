package wdk

import (
	"sync"
)

// ItemConfirmInfo 结构体
type ItemConfirmInfo struct {
	// 确认数量(为正数或零)
	ConfirmQuantity string `json:"confirm_quantity,omitempty" xml:"confirm_quantity,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}

var poolItemConfirmInfo = sync.Pool{
	New: func() any {
		return new(ItemConfirmInfo)
	},
}

// GetItemConfirmInfo() 从对象池中获取ItemConfirmInfo
func GetItemConfirmInfo() *ItemConfirmInfo {
	return poolItemConfirmInfo.Get().(*ItemConfirmInfo)
}

// ReleaseItemConfirmInfo 释放ItemConfirmInfo
func ReleaseItemConfirmInfo(v *ItemConfirmInfo) {
	v.ConfirmQuantity = ""
	v.SkuCode = ""
	poolItemConfirmInfo.Put(v)
}
