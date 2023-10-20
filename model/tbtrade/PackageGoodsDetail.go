package tbtrade

import (
	"sync"
)

// PackageGoodsDetail 结构体
type PackageGoodsDetail struct {
	// 商品数字编号
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku_id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolPackageGoodsDetail = sync.Pool{
	New: func() any {
		return new(PackageGoodsDetail)
	},
}

// GetPackageGoodsDetail() 从对象池中获取PackageGoodsDetail
func GetPackageGoodsDetail() *PackageGoodsDetail {
	return poolPackageGoodsDetail.Get().(*PackageGoodsDetail)
}

// ReleasePackageGoodsDetail 释放PackageGoodsDetail
func ReleasePackageGoodsDetail(v *PackageGoodsDetail) {
	v.ItemId = 0
	v.SkuId = 0
	v.Amount = 0
	poolPackageGoodsDetail.Put(v)
}
