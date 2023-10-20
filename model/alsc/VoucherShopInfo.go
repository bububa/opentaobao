package alsc

import (
	"sync"
)

// VoucherShopInfo 结构体
type VoucherShopInfo struct {
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolVoucherShopInfo = sync.Pool{
	New: func() any {
		return new(VoucherShopInfo)
	},
}

// GetVoucherShopInfo() 从对象池中获取VoucherShopInfo
func GetVoucherShopInfo() *VoucherShopInfo {
	return poolVoucherShopInfo.Get().(*VoucherShopInfo)
}

// ReleaseVoucherShopInfo 释放VoucherShopInfo
func ReleaseVoucherShopInfo(v *VoucherShopInfo) {
	v.OutShopId = ""
	v.ShopId = ""
	poolVoucherShopInfo.Put(v)
}
