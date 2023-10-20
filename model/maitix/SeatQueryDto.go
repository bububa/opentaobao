package maitix

import (
	"sync"
)

// SeatQueryDto 结构体
type SeatQueryDto struct {
	// 单票座位信息
	OrdinarySeatDTOS []OrdinarySeatDto `json:"ordinary_seat_d_t_o_s,omitempty" xml:"ordinary_seat_d_t_o_s>ordinary_seat_dto,omitempty"`
	// 套票座位信息
	CombineSeatDTOS []CombineSeatDto `json:"combine_seat_d_t_o_s,omitempty" xml:"combine_seat_d_t_o_s>combine_seat_dto,omitempty"`
	// 座位状态，2有效
	SeatStatus int64 `json:"seat_status,omitempty" xml:"seat_status,omitempty"`
}

var poolSeatQueryDto = sync.Pool{
	New: func() any {
		return new(SeatQueryDto)
	},
}

// GetSeatQueryDto() 从对象池中获取SeatQueryDto
func GetSeatQueryDto() *SeatQueryDto {
	return poolSeatQueryDto.Get().(*SeatQueryDto)
}

// ReleaseSeatQueryDto 释放SeatQueryDto
func ReleaseSeatQueryDto(v *SeatQueryDto) {
	v.OrdinarySeatDTOS = v.OrdinarySeatDTOS[:0]
	v.CombineSeatDTOS = v.CombineSeatDTOS[:0]
	v.SeatStatus = 0
	poolSeatQueryDto.Put(v)
}
