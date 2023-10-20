package kbalgo

import (
	"sync"
)

// ShopProduct 结构体
type ShopProduct struct {
	// Product_info
	ProductInfos []ProductInfo `json:"product_infos,omitempty" xml:"product_infos>product_info,omitempty"`
	// 优惠券信息
	CouponInfos []CouponInfo `json:"coupon_infos,omitempty" xml:"coupon_infos>coupon_info,omitempty"`
}

var poolShopProduct = sync.Pool{
	New: func() any {
		return new(ShopProduct)
	},
}

// GetShopProduct() 从对象池中获取ShopProduct
func GetShopProduct() *ShopProduct {
	return poolShopProduct.Get().(*ShopProduct)
}

// ReleaseShopProduct 释放ShopProduct
func ReleaseShopProduct(v *ShopProduct) {
	v.ProductInfos = v.ProductInfos[:0]
	v.CouponInfos = v.CouponInfos[:0]
	poolShopProduct.Put(v)
}
