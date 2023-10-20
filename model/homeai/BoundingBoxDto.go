package homeai

import (
	"sync"
)

// BoundingBoxDto 结构体
type BoundingBoxDto struct {
	// x
	XLen string `json:"x_len,omitempty" xml:"x_len,omitempty"`
	// y
	YLen string `json:"y_len,omitempty" xml:"y_len,omitempty"`
	// z
	ZLen string `json:"z_len,omitempty" xml:"z_len,omitempty"`
}

var poolBoundingBoxDto = sync.Pool{
	New: func() any {
		return new(BoundingBoxDto)
	},
}

// GetBoundingBoxDto() 从对象池中获取BoundingBoxDto
func GetBoundingBoxDto() *BoundingBoxDto {
	return poolBoundingBoxDto.Get().(*BoundingBoxDto)
}

// ReleaseBoundingBoxDto 释放BoundingBoxDto
func ReleaseBoundingBoxDto(v *BoundingBoxDto) {
	v.XLen = ""
	v.YLen = ""
	v.ZLen = ""
	poolBoundingBoxDto.Put(v)
}
