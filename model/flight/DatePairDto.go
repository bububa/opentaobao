package flight

import (
	"sync"
)

// DatePairDto 结构体
type DatePairDto struct {
	// 允许航班截止日期，无需传入时分秒
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 允许航班起始日期，无需传入时分秒
	Start string `json:"start,omitempty" xml:"start,omitempty"`
}

var poolDatePairDto = sync.Pool{
	New: func() any {
		return new(DatePairDto)
	},
}

// GetDatePairDto() 从对象池中获取DatePairDto
func GetDatePairDto() *DatePairDto {
	return poolDatePairDto.Get().(*DatePairDto)
}

// ReleaseDatePairDto 释放DatePairDto
func ReleaseDatePairDto(v *DatePairDto) {
	v.End = ""
	v.Start = ""
	poolDatePairDto.Put(v)
}
