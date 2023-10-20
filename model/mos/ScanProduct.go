package mos

import (
	"sync"
)

// ScanProduct 结构体
type ScanProduct struct {
	// 销售属性
	SalePropertys []SaleProperty `json:"sale_propertys,omitempty" xml:"sale_propertys>sale_property,omitempty"`
	// 货号
	ArtNo string `json:"art_no,omitempty" xml:"art_no,omitempty"`
	// 条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 集团专柜编码
	GroupShopCode string `json:"group_shop_code,omitempty" xml:"group_shop_code,omitempty"`
	// 收银编码
	IntimeCodes string `json:"intime_codes,omitempty" xml:"intime_codes,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 销售价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 专柜Code
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 商品Id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品来源
	SourceType string `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 商品来源
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 吊牌价
	TagPrice string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
	// 默认收银编码
	DefaultIntimeCode string `json:"default_intime_code,omitempty" xml:"default_intime_code,omitempty"`
	// 唯一码
	UniqueCode string `json:"unique_code,omitempty" xml:"unique_code,omitempty"`
	// 商品标签
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
}

var poolScanProduct = sync.Pool{
	New: func() any {
		return new(ScanProduct)
	},
}

// GetScanProduct() 从对象池中获取ScanProduct
func GetScanProduct() *ScanProduct {
	return poolScanProduct.Get().(*ScanProduct)
}

// ReleaseScanProduct 释放ScanProduct
func ReleaseScanProduct(v *ScanProduct) {
	v.SalePropertys = v.SalePropertys[:0]
	v.ArtNo = ""
	v.BarCode = ""
	v.GroupShopCode = ""
	v.IntimeCodes = ""
	v.Name = ""
	v.Price = ""
	v.ShopCode = ""
	v.SkuId = ""
	v.SourceType = ""
	v.StoreCode = ""
	v.TagPrice = ""
	v.DefaultIntimeCode = ""
	v.UniqueCode = ""
	v.Tag = ""
	poolScanProduct.Put(v)
}
