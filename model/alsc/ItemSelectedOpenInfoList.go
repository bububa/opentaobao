package alsc

import (
	"sync"
)

// ItemSelectedOpenInfoList 结构体
type ItemSelectedOpenInfoList struct {
	// 规格id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 菜单id
	DishId string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
	// 规则号
	SkuOutNo string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
	// 菜品号
	DishOutNo string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
}

var poolItemSelectedOpenInfoList = sync.Pool{
	New: func() any {
		return new(ItemSelectedOpenInfoList)
	},
}

// GetItemSelectedOpenInfoList() 从对象池中获取ItemSelectedOpenInfoList
func GetItemSelectedOpenInfoList() *ItemSelectedOpenInfoList {
	return poolItemSelectedOpenInfoList.Get().(*ItemSelectedOpenInfoList)
}

// ReleaseItemSelectedOpenInfoList 释放ItemSelectedOpenInfoList
func ReleaseItemSelectedOpenInfoList(v *ItemSelectedOpenInfoList) {
	v.SkuId = ""
	v.DishId = ""
	v.SkuOutNo = ""
	v.DishOutNo = ""
	poolItemSelectedOpenInfoList.Put(v)
}
