package alsc

import (
	"sync"
)

// VoucherItemInfo 结构体
type VoucherItemInfo struct {
	// sku规格id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 菜品id
	DishId string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
	// 外部sku规格id
	SkuOutNo string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
	// 外部菜品id
	DishOutNo string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
}

var poolVoucherItemInfo = sync.Pool{
	New: func() any {
		return new(VoucherItemInfo)
	},
}

// GetVoucherItemInfo() 从对象池中获取VoucherItemInfo
func GetVoucherItemInfo() *VoucherItemInfo {
	return poolVoucherItemInfo.Get().(*VoucherItemInfo)
}

// ReleaseVoucherItemInfo 释放VoucherItemInfo
func ReleaseVoucherItemInfo(v *VoucherItemInfo) {
	v.SkuId = ""
	v.DishId = ""
	v.SkuOutNo = ""
	v.DishOutNo = ""
	poolVoucherItemInfo.Put(v)
}
