package alihealth2

import (
	"sync"
)

// TopIsvStqtyLstDto 结构体
type TopIsvStqtyLstDto struct {
	// 门店编码
	DeptCode string `json:"dept_code,omitempty" xml:"dept_code,omitempty"`
	// 商品编码
	GoodsCode string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
	// 数量
	GoodsQty string `json:"goods_qty,omitempty" xml:"goods_qty,omitempty"`
}

var poolTopIsvStqtyLstDto = sync.Pool{
	New: func() any {
		return new(TopIsvStqtyLstDto)
	},
}

// GetTopIsvStqtyLstDto() 从对象池中获取TopIsvStqtyLstDto
func GetTopIsvStqtyLstDto() *TopIsvStqtyLstDto {
	return poolTopIsvStqtyLstDto.Get().(*TopIsvStqtyLstDto)
}

// ReleaseTopIsvStqtyLstDto 释放TopIsvStqtyLstDto
func ReleaseTopIsvStqtyLstDto(v *TopIsvStqtyLstDto) {
	v.DeptCode = ""
	v.GoodsCode = ""
	v.GoodsQty = ""
	poolTopIsvStqtyLstDto.Put(v)
}
