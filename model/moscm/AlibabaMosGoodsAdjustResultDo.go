package moscm

import (
	"sync"
)

// AlibabaMosGoodsAdjustResultDo 结构体
type AlibabaMosGoodsAdjustResultDo struct {
	// 返回生成的单据号
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaMosGoodsAdjustResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsAdjustResultDo)
	},
}

// GetAlibabaMosGoodsAdjustResultDo() 从对象池中获取AlibabaMosGoodsAdjustResultDo
func GetAlibabaMosGoodsAdjustResultDo() *AlibabaMosGoodsAdjustResultDo {
	return poolAlibabaMosGoodsAdjustResultDo.Get().(*AlibabaMosGoodsAdjustResultDo)
}

// ReleaseAlibabaMosGoodsAdjustResultDo 释放AlibabaMosGoodsAdjustResultDo
func ReleaseAlibabaMosGoodsAdjustResultDo(v *AlibabaMosGoodsAdjustResultDo) {
	v.Data = ""
	poolAlibabaMosGoodsAdjustResultDo.Put(v)
}
