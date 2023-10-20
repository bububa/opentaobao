package tbk

import (
	"sync"
)

// TopNInfoDto 结构体
type TopNInfoDto struct {
	// 前N件佣金结束时间
	TopnEndTime string `json:"topn_end_time,omitempty" xml:"topn_end_time,omitempty"`
	// 前N件佣金开始时间
	TopnStartTime string `json:"topn_start_time,omitempty" xml:"topn_start_time,omitempty"`
	// 前N件佣金率
	TopnRate string `json:"topn_rate,omitempty" xml:"topn_rate,omitempty"`
	// 前N件剩余库存
	TopnQuantity int64 `json:"topn_quantity,omitempty" xml:"topn_quantity,omitempty"`
	// 前N件初始总库存
	TopnTotalCount int64 `json:"topn_total_count,omitempty" xml:"topn_total_count,omitempty"`
}

var poolTopNInfoDto = sync.Pool{
	New: func() any {
		return new(TopNInfoDto)
	},
}

// GetTopNInfoDto() 从对象池中获取TopNInfoDto
func GetTopNInfoDto() *TopNInfoDto {
	return poolTopNInfoDto.Get().(*TopNInfoDto)
}

// ReleaseTopNInfoDto 释放TopNInfoDto
func ReleaseTopNInfoDto(v *TopNInfoDto) {
	v.TopnEndTime = ""
	v.TopnStartTime = ""
	v.TopnRate = ""
	v.TopnQuantity = 0
	v.TopnTotalCount = 0
	poolTopNInfoDto.Put(v)
}
