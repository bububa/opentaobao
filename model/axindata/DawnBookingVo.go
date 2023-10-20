package axindata

import (
	"sync"
)

// DawnBookingVo 结构体
type DawnBookingVo struct {
	// 凌晨房预订开始时间
	DawnBookStartTime string `json:"dawn_book_start_time,omitempty" xml:"dawn_book_start_time,omitempty"`
	// 凌晨房预订结束时间
	DawnBookEndTime string `json:"dawn_book_end_time,omitempty" xml:"dawn_book_end_time,omitempty"`
	// 是否支持凌晨房
	CanDawnBook bool `json:"can_dawn_book,omitempty" xml:"can_dawn_book,omitempty"`
}

var poolDawnBookingVo = sync.Pool{
	New: func() any {
		return new(DawnBookingVo)
	},
}

// GetDawnBookingVo() 从对象池中获取DawnBookingVo
func GetDawnBookingVo() *DawnBookingVo {
	return poolDawnBookingVo.Get().(*DawnBookingVo)
}

// ReleaseDawnBookingVo 释放DawnBookingVo
func ReleaseDawnBookingVo(v *DawnBookingVo) {
	v.DawnBookStartTime = ""
	v.DawnBookEndTime = ""
	v.CanDawnBook = false
	poolDawnBookingVo.Put(v)
}
