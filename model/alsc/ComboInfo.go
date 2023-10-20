package alsc

import (
	"sync"
)

// ComboInfo 结构体
type ComboInfo struct {
	// 套餐菜品做法明细
	CookingMethodsInfoList []AttachInfo `json:"cooking_methods_info_list,omitempty" xml:"cooking_methods_info_list>attach_info,omitempty"`
	// 套餐菜品配料明细
	IngredientsInfoList []AttachInfo `json:"ingredients_info_list,omitempty" xml:"ingredients_info_list>attach_info,omitempty"`
	// 商品id
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 商品名称
	OutItemName string `json:"out_item_name,omitempty" xml:"out_item_name,omitempty"`
	// 商品规格id
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 商品规格名称
	OutSkuName string `json:"out_sku_name,omitempty" xml:"out_sku_name,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 商品数量
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
	// 单价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolComboInfo = sync.Pool{
	New: func() any {
		return new(ComboInfo)
	},
}

// GetComboInfo() 从对象池中获取ComboInfo
func GetComboInfo() *ComboInfo {
	return poolComboInfo.Get().(*ComboInfo)
}

// ReleaseComboInfo 释放ComboInfo
func ReleaseComboInfo(v *ComboInfo) {
	v.CookingMethodsInfoList = v.CookingMethodsInfoList[:0]
	v.IngredientsInfoList = v.IngredientsInfoList[:0]
	v.OutItemId = ""
	v.OutItemName = ""
	v.OutSkuId = ""
	v.OutSkuName = ""
	v.Unit = ""
	v.Weight = ""
	v.ItemCount = 0
	v.Price = 0
	poolComboInfo.Put(v)
}
