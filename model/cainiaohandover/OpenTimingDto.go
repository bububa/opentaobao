package cainiaohandover

import (
	"sync"
)

// OpenTimingDto 结构体
type OpenTimingDto struct {
	// 展示文案
	DisplayText string `json:"display_text,omitempty" xml:"display_text,omitempty"`
	// 时效类型，ESTIMATE：预估时效，PROMISE：承诺时效
	TimingType string `json:"timing_type,omitempty" xml:"timing_type,omitempty"`
	// 最慢时效
	SlowestTiming int64 `json:"slowest_timing,omitempty" xml:"slowest_timing,omitempty"`
	// 最快时效
	FastTiming int64 `json:"fast_timing,omitempty" xml:"fast_timing,omitempty"`
}

var poolOpenTimingDto = sync.Pool{
	New: func() any {
		return new(OpenTimingDto)
	},
}

// GetOpenTimingDto() 从对象池中获取OpenTimingDto
func GetOpenTimingDto() *OpenTimingDto {
	return poolOpenTimingDto.Get().(*OpenTimingDto)
}

// ReleaseOpenTimingDto 释放OpenTimingDto
func ReleaseOpenTimingDto(v *OpenTimingDto) {
	v.DisplayText = ""
	v.TimingType = ""
	v.SlowestTiming = 0
	v.FastTiming = 0
	poolOpenTimingDto.Put(v)
}
