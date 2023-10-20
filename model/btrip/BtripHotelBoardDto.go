package btrip

import (
	"sync"
)

// BtripHotelBoardDto 结构体
type BtripHotelBoardDto struct {
	// 餐食数量
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
	// 餐食类型
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
}

var poolBtripHotelBoardDto = sync.Pool{
	New: func() any {
		return new(BtripHotelBoardDto)
	},
}

// GetBtripHotelBoardDto() 从对象池中获取BtripHotelBoardDto
func GetBtripHotelBoardDto() *BtripHotelBoardDto {
	return poolBtripHotelBoardDto.Get().(*BtripHotelBoardDto)
}

// ReleaseBtripHotelBoardDto 释放BtripHotelBoardDto
func ReleaseBtripHotelBoardDto(v *BtripHotelBoardDto) {
	v.BoardNum = 0
	v.BoardType = 0
	poolBtripHotelBoardDto.Put(v)
}
