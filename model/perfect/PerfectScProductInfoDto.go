package perfect

import (
	"sync"
)

// PerfectScProductInfoDto 结构体
type PerfectScProductInfoDto struct {
	// 货品编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 货品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolPerfectScProductInfoDto = sync.Pool{
	New: func() any {
		return new(PerfectScProductInfoDto)
	},
}

// GetPerfectScProductInfoDto() 从对象池中获取PerfectScProductInfoDto
func GetPerfectScProductInfoDto() *PerfectScProductInfoDto {
	return poolPerfectScProductInfoDto.Get().(*PerfectScProductInfoDto)
}

// ReleasePerfectScProductInfoDto 释放PerfectScProductInfoDto
func ReleasePerfectScProductInfoDto(v *PerfectScProductInfoDto) {
	v.OuterId = ""
	v.ProductId = 0
	poolPerfectScProductInfoDto.Put(v)
}
