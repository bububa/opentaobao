package feedflow

import (
	"sync"
)

// TimeSpanDto 结构体
type TimeSpanDto struct {
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 折扣率
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolTimeSpanDto = sync.Pool{
	New: func() any {
		return new(TimeSpanDto)
	},
}

// GetTimeSpanDto() 从对象池中获取TimeSpanDto
func GetTimeSpanDto() *TimeSpanDto {
	return poolTimeSpanDto.Get().(*TimeSpanDto)
}

// ReleaseTimeSpanDto 释放TimeSpanDto
func ReleaseTimeSpanDto(v *TimeSpanDto) {
	v.Time = ""
	v.Discount = 0
	poolTimeSpanDto.Put(v)
}
