package alsc

import (
	"sync"
)

// ItemSelectedOpenInfo 结构体
type ItemSelectedOpenInfo struct {
	// 规则ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 菜品ID
	DishId string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
	// 外部菜品ID
	DishOutNo string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
	// 外部规则ID
	SkuOutNo string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
}

var poolItemSelectedOpenInfo = sync.Pool{
	New: func() any {
		return new(ItemSelectedOpenInfo)
	},
}

// GetItemSelectedOpenInfo() 从对象池中获取ItemSelectedOpenInfo
func GetItemSelectedOpenInfo() *ItemSelectedOpenInfo {
	return poolItemSelectedOpenInfo.Get().(*ItemSelectedOpenInfo)
}

// ReleaseItemSelectedOpenInfo 释放ItemSelectedOpenInfo
func ReleaseItemSelectedOpenInfo(v *ItemSelectedOpenInfo) {
	v.SkuId = ""
	v.DishId = ""
	v.DishOutNo = ""
	v.SkuOutNo = ""
	poolItemSelectedOpenInfo.Put(v)
}
