package alihouse

import (
	"sync"
)

// GoodsLayoutDto 结构体
type GoodsLayoutDto struct {
	// 户型所属公寓外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 户型外部id
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 户型货外部id
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 户型业务类型
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolGoodsLayoutDto = sync.Pool{
	New: func() any {
		return new(GoodsLayoutDto)
	},
}

// GetGoodsLayoutDto() 从对象池中获取GoodsLayoutDto
func GetGoodsLayoutDto() *GoodsLayoutDto {
	return poolGoodsLayoutDto.Get().(*GoodsLayoutDto)
}

// ReleaseGoodsLayoutDto 释放GoodsLayoutDto
func ReleaseGoodsLayoutDto(v *GoodsLayoutDto) {
	v.OuterId = ""
	v.OuterLayoutId = ""
	v.OuterTid = ""
	v.BusinessType = 0
	poolGoodsLayoutDto.Put(v)
}
