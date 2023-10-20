package product

import (
	"sync"
)

// ItemSkuStatus 结构体
type ItemSkuStatus struct {
	// sku集合
	SkuStatusList []SkuStatus `json:"sku_status_list,omitempty" xml:"sku_status_list>sku_status,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolItemSkuStatus = sync.Pool{
	New: func() any {
		return new(ItemSkuStatus)
	},
}

// GetItemSkuStatus() 从对象池中获取ItemSkuStatus
func GetItemSkuStatus() *ItemSkuStatus {
	return poolItemSkuStatus.Get().(*ItemSkuStatus)
}

// ReleaseItemSkuStatus 释放ItemSkuStatus
func ReleaseItemSkuStatus(v *ItemSkuStatus) {
	v.SkuStatusList = v.SkuStatusList[:0]
	v.Title = ""
	v.Status = 0
	poolItemSkuStatus.Put(v)
}
