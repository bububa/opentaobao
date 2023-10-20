package ascpchannel

import (
	"sync"
)

// ProductDetailQueryResponseForSupplier 结构体
type ProductDetailQueryResponseForSupplier struct {
	// sku 列表
	SkuList []SkuList `json:"sku_list,omitempty" xml:"sku_list>sku_list,omitempty"`
	// 经营模式,agent:代销，dealer经销
	SalesModeList []string `json:"sales_mode_list,omitempty" xml:"sales_mode_list>string,omitempty"`
	// 图片链接列表
	PictureList []string `json:"picture_list,omitempty" xml:"picture_list>string,omitempty"`
	// 销售属性，格式[k1:v2,k2:v2]
	PropertyList []string `json:"property_list,omitempty" xml:"property_list>string,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道产品名称
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 商家编码
	OutNo string `json:"out_no,omitempty" xml:"out_no,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 产品 id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品货品 id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolProductDetailQueryResponseForSupplier = sync.Pool{
	New: func() any {
		return new(ProductDetailQueryResponseForSupplier)
	},
}

// GetProductDetailQueryResponseForSupplier() 从对象池中获取ProductDetailQueryResponseForSupplier
func GetProductDetailQueryResponseForSupplier() *ProductDetailQueryResponseForSupplier {
	return poolProductDetailQueryResponseForSupplier.Get().(*ProductDetailQueryResponseForSupplier)
}

// ReleaseProductDetailQueryResponseForSupplier 释放ProductDetailQueryResponseForSupplier
func ReleaseProductDetailQueryResponseForSupplier(v *ProductDetailQueryResponseForSupplier) {
	v.SkuList = v.SkuList[:0]
	v.SalesModeList = v.SalesModeList[:0]
	v.PictureList = v.PictureList[:0]
	v.PropertyList = v.PropertyList[:0]
	v.SubChannelCode = ""
	v.ProductTitle = ""
	v.Category = ""
	v.OutNo = ""
	v.Brand = ""
	v.ChannelCode = ""
	v.ProductId = 0
	v.ScItemId = 0
	poolProductDetailQueryResponseForSupplier.Put(v)
}
