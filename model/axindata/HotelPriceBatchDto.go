package axindata

import (
	"sync"
)

// HotelPriceBatchDto 结构体
type HotelPriceBatchDto struct {
	// 房型列表，里面包含了rate信息列有
	RoomList []RoomPriceDto `json:"room_list,omitempty" xml:"room_list>room_price_dto,omitempty"`
	// 当前酒店查询异常code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 当前酒店查询异常信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 当前酒店查询是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHotelPriceBatchDto = sync.Pool{
	New: func() any {
		return new(HotelPriceBatchDto)
	},
}

// GetHotelPriceBatchDto() 从对象池中获取HotelPriceBatchDto
func GetHotelPriceBatchDto() *HotelPriceBatchDto {
	return poolHotelPriceBatchDto.Get().(*HotelPriceBatchDto)
}

// ReleaseHotelPriceBatchDto 释放HotelPriceBatchDto
func ReleaseHotelPriceBatchDto(v *HotelPriceBatchDto) {
	v.RoomList = v.RoomList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Shid = 0
	v.Success = false
	poolHotelPriceBatchDto.Put(v)
}
