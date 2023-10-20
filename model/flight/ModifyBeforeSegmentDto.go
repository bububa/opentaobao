package flight

import (
	"sync"
)

// ModifyBeforeSegmentDto 结构体
type ModifyBeforeSegmentDto struct {
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 起飞城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
}

var poolModifyBeforeSegmentDto = sync.Pool{
	New: func() any {
		return new(ModifyBeforeSegmentDto)
	},
}

// GetModifyBeforeSegmentDto() 从对象池中获取ModifyBeforeSegmentDto
func GetModifyBeforeSegmentDto() *ModifyBeforeSegmentDto {
	return poolModifyBeforeSegmentDto.Get().(*ModifyBeforeSegmentDto)
}

// ReleaseModifyBeforeSegmentDto 释放ModifyBeforeSegmentDto
func ReleaseModifyBeforeSegmentDto(v *ModifyBeforeSegmentDto) {
	v.ArrCity = ""
	v.DepCity = ""
	poolModifyBeforeSegmentDto.Put(v)
}
