package idleisv

import (
	"sync"
)

// IdleItemDistributionDo 结构体
type IdleItemDistributionDo struct {
	// 货品供货价，国际分销场景中使用
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 后端货品id,在国际分销场景下使用
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 对应的货品库存等级，在库良品：1，临期品：139，一级残次：137，二级残次：138，在库残次：101
	InvGrade string `json:"inv_grade,omitempty" xml:"inv_grade,omitempty"`
}

var poolIdleItemDistributionDo = sync.Pool{
	New: func() any {
		return new(IdleItemDistributionDo)
	},
}

// GetIdleItemDistributionDo() 从对象池中获取IdleItemDistributionDo
func GetIdleItemDistributionDo() *IdleItemDistributionDo {
	return poolIdleItemDistributionDo.Get().(*IdleItemDistributionDo)
}

// ReleaseIdleItemDistributionDo 释放IdleItemDistributionDo
func ReleaseIdleItemDistributionDo(v *IdleItemDistributionDo) {
	v.ProductPrice = ""
	v.ProductId = ""
	v.InvGrade = ""
	poolIdleItemDistributionDo.Put(v)
}
