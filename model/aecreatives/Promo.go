package aecreatives

import (
	"sync"
)

// Promo 结构体
type Promo struct {
	// 主题活动描述
	PromoDesc string `json:"promo_desc,omitempty" xml:"promo_desc,omitempty"`
	// 主题活动名称
	PromoName string `json:"promo_name,omitempty" xml:"promo_name,omitempty"`
	// 主题活动的商品数量
	ProductNum int64 `json:"product_num,omitempty" xml:"product_num,omitempty"`
}

var poolPromo = sync.Pool{
	New: func() any {
		return new(Promo)
	},
}

// GetPromo() 从对象池中获取Promo
func GetPromo() *Promo {
	return poolPromo.Get().(*Promo)
}

// ReleasePromo 释放Promo
func ReleasePromo(v *Promo) {
	v.PromoDesc = ""
	v.PromoName = ""
	v.ProductNum = 0
	poolPromo.Put(v)
}
