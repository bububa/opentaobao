package wdk

import (
	"sync"
)

// SkuFeatureDo 结构体
type SkuFeatureDo struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 门店编码，用来给特定门店商品标记
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 机构编码, 可以指定机构的商品标记，如果要全量商品请填写商家编码
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 渠道编码，
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 标记编码，需要和调用方约定
	FeatureCode string `json:"feature_code,omitempty" xml:"feature_code,omitempty"`
	// 标记值，需要和调用方约定
	FeatureValue string `json:"feature_value,omitempty" xml:"feature_value,omitempty"`
	// 是否添加，默认是添加
	Add bool `json:"add,omitempty" xml:"add,omitempty"`
}

var poolSkuFeatureDo = sync.Pool{
	New: func() any {
		return new(SkuFeatureDo)
	},
}

// GetSkuFeatureDo() 从对象池中获取SkuFeatureDo
func GetSkuFeatureDo() *SkuFeatureDo {
	return poolSkuFeatureDo.Get().(*SkuFeatureDo)
}

// ReleaseSkuFeatureDo 释放SkuFeatureDo
func ReleaseSkuFeatureDo(v *SkuFeatureDo) {
	v.SkuCode = ""
	v.OuCode = ""
	v.OrgCode = ""
	v.ChannelCode = ""
	v.FeatureCode = ""
	v.FeatureValue = ""
	v.Add = false
	poolSkuFeatureDo.Put(v)
}
