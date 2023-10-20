package openmall

import (
	"sync"
)

// PostDo 结构体
type PostDo struct {
	// 运费金额，运费0为包邮
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 发货方式 快递，EMS等
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 物流方式，可选值  ems, post, express
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
}

var poolPostDo = sync.Pool{
	New: func() any {
		return new(PostDo)
	},
}

// GetPostDo() 从对象池中获取PostDo
func GetPostDo() *PostDo {
	return poolPostDo.Get().(*PostDo)
}

// ReleasePostDo 释放PostDo
func ReleasePostDo(v *PostDo) {
	v.Amount = ""
	v.Name = ""
	v.ShippingType = ""
	poolPostDo.Put(v)
}
