package scbp

import (
	"sync"
)

// TopContextDto 结构体
type TopContextDto struct {
	// 产品线id
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolTopContextDto = sync.Pool{
	New: func() any {
		return new(TopContextDto)
	},
}

// GetTopContextDto() 从对象池中获取TopContextDto
func GetTopContextDto() *TopContextDto {
	return poolTopContextDto.Get().(*TopContextDto)
}

// ReleaseTopContextDto 释放TopContextDto
func ReleaseTopContextDto(v *TopContextDto) {
	v.ProductLineId = 0
	v.ProductId = 0
	poolTopContextDto.Put(v)
}
