package scbp

import (
	"sync"
)

// AdProductEffectDto 结构体
type AdProductEffectDto struct {
	// 曝光
	ImprCnt int64 `json:"impr_cnt,omitempty" xml:"impr_cnt,omitempty"`
	// 点击
	ClickCnt int64 `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 消耗(分)
	CostAmt int64 `json:"cost_amt,omitempty" xml:"cost_amt,omitempty"`
}

var poolAdProductEffectDto = sync.Pool{
	New: func() any {
		return new(AdProductEffectDto)
	},
}

// GetAdProductEffectDto() 从对象池中获取AdProductEffectDto
func GetAdProductEffectDto() *AdProductEffectDto {
	return poolAdProductEffectDto.Get().(*AdProductEffectDto)
}

// ReleaseAdProductEffectDto 释放AdProductEffectDto
func ReleaseAdProductEffectDto(v *AdProductEffectDto) {
	v.ImprCnt = 0
	v.ClickCnt = 0
	v.CostAmt = 0
	poolAdProductEffectDto.Put(v)
}
