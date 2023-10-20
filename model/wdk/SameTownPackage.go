package wdk

import (
	"sync"
)

// SameTownPackage 结构体
type SameTownPackage struct {
	// 令牌号
	TokenCode string `json:"token_code,omitempty" xml:"token_code,omitempty"`
	// 6位交接码
	PickupCode string `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`
	// 包裹中商品出库销售数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 包裹中商品出库库存数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// 周转箱
	Container *Container `json:"container,omitempty" xml:"container,omitempty"`
}

var poolSameTownPackage = sync.Pool{
	New: func() any {
		return new(SameTownPackage)
	},
}

// GetSameTownPackage() 从对象池中获取SameTownPackage
func GetSameTownPackage() *SameTownPackage {
	return poolSameTownPackage.Get().(*SameTownPackage)
}

// ReleaseSameTownPackage 释放SameTownPackage
func ReleaseSameTownPackage(v *SameTownPackage) {
	v.TokenCode = ""
	v.PickupCode = ""
	v.ActualSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.Container = nil
	poolSameTownPackage.Put(v)
}
