package wdk

import (
	"sync"
)

// IsvShopStockParam 结构体
type IsvShopStockParam struct {
	// 门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 派样活动id
	SampleActivityId int64 `json:"sample_activity_id,omitempty" xml:"sample_activity_id,omitempty"`
}

var poolIsvShopStockParam = sync.Pool{
	New: func() any {
		return new(IsvShopStockParam)
	},
}

// GetIsvShopStockParam() 从对象池中获取IsvShopStockParam
func GetIsvShopStockParam() *IsvShopStockParam {
	return poolIsvShopStockParam.Get().(*IsvShopStockParam)
}

// ReleaseIsvShopStockParam 释放IsvShopStockParam
func ReleaseIsvShopStockParam(v *IsvShopStockParam) {
	v.ShopCode = ""
	v.MerchantCode = ""
	v.Barcode = ""
	v.SampleActivityId = 0
	poolIsvShopStockParam.Put(v)
}
