package idleisv

import (
	"sync"
)

// SpuPVDo 结构体
type SpuPVDo struct {
	// 品牌和型号信息
	BrandModelList []IdleNewPubValueDo `json:"brand_model_list,omitempty" xml:"brand_model_list>idle_new_pub_value_do,omitempty"`
	// 品牌型号的显示值
	SpuTitle string `json:"spu_title,omitempty" xml:"spu_title,omitempty"`
}

var poolSpuPVDo = sync.Pool{
	New: func() any {
		return new(SpuPVDo)
	},
}

// GetSpuPVDo() 从对象池中获取SpuPVDo
func GetSpuPVDo() *SpuPVDo {
	return poolSpuPVDo.Get().(*SpuPVDo)
}

// ReleaseSpuPVDo 释放SpuPVDo
func ReleaseSpuPVDo(v *SpuPVDo) {
	v.BrandModelList = v.BrandModelList[:0]
	v.SpuTitle = ""
	poolSpuPVDo.Put(v)
}
