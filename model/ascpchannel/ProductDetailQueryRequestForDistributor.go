package ascpchannel

import (
	"sync"
)

// ProductDetailQueryRequestForDistributor 结构体
type ProductDetailQueryRequestForDistributor struct {
	// 产品 id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 是否查询sku信息
	IncludeSku bool `json:"include_sku,omitempty" xml:"include_sku,omitempty"`
}

var poolProductDetailQueryRequestForDistributor = sync.Pool{
	New: func() any {
		return new(ProductDetailQueryRequestForDistributor)
	},
}

// GetProductDetailQueryRequestForDistributor() 从对象池中获取ProductDetailQueryRequestForDistributor
func GetProductDetailQueryRequestForDistributor() *ProductDetailQueryRequestForDistributor {
	return poolProductDetailQueryRequestForDistributor.Get().(*ProductDetailQueryRequestForDistributor)
}

// ReleaseProductDetailQueryRequestForDistributor 释放ProductDetailQueryRequestForDistributor
func ReleaseProductDetailQueryRequestForDistributor(v *ProductDetailQueryRequestForDistributor) {
	v.ProductId = ""
	v.SubChannelCode = ""
	v.ChannelCode = ""
	v.MerchantCode = ""
	v.IncludeSku = false
	poolProductDetailQueryRequestForDistributor.Put(v)
}
