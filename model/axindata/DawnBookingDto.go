package axindata

import (
	"sync"
)

// DawnBookingDto 结构体
type DawnBookingDto struct {
	// 凌晨房预订开始时间
	DawnBookStartTime string `json:"dawn_book_start_time,omitempty" xml:"dawn_book_start_time,omitempty"`
	// 凌晨房预订结束时间
	DawnBookEndTime string `json:"dawn_book_end_time,omitempty" xml:"dawn_book_end_time,omitempty"`
	// 是否支持凌晨房
	CanDawnBook bool `json:"can_dawn_book,omitempty" xml:"can_dawn_book,omitempty"`
}

var poolDawnBookingDto = sync.Pool{
	New: func() any {
		return new(DawnBookingDto)
	},
}

// GetDawnBookingDto() 从对象池中获取DawnBookingDto
func GetDawnBookingDto() *DawnBookingDto {
	return poolDawnBookingDto.Get().(*DawnBookingDto)
}

// ReleaseDawnBookingDto 释放DawnBookingDto
func ReleaseDawnBookingDto(v *DawnBookingDto) {
	v.DawnBookStartTime = ""
	v.DawnBookEndTime = ""
	v.CanDawnBook = false
	poolDawnBookingDto.Put(v)
}
