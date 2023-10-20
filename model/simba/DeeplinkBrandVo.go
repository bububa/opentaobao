package simba

import (
	"sync"
)

// DeeplinkBrandVo 结构体
type DeeplinkBrandVo struct {
	// 品牌id
	DeeplinkBrandId string `json:"deeplink_brand_id,omitempty" xml:"deeplink_brand_id,omitempty"`
	// 品牌名字
	DeeplinkBrandName string `json:"deeplink_brand_name,omitempty" xml:"deeplink_brand_name,omitempty"`
}

var poolDeeplinkBrandVo = sync.Pool{
	New: func() any {
		return new(DeeplinkBrandVo)
	},
}

// GetDeeplinkBrandVo() 从对象池中获取DeeplinkBrandVo
func GetDeeplinkBrandVo() *DeeplinkBrandVo {
	return poolDeeplinkBrandVo.Get().(*DeeplinkBrandVo)
}

// ReleaseDeeplinkBrandVo 释放DeeplinkBrandVo
func ReleaseDeeplinkBrandVo(v *DeeplinkBrandVo) {
	v.DeeplinkBrandId = ""
	v.DeeplinkBrandName = ""
	poolDeeplinkBrandVo.Put(v)
}
