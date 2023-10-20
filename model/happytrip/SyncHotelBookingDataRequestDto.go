package happytrip

import (
	"sync"
)

// SyncHotelBookingDataRequestDto 结构体
type SyncHotelBookingDataRequestDto struct {
	// 预订信息数据
	Segments []string `json:"segments,omitempty" xml:"segments>string,omitempty"`
}

var poolSyncHotelBookingDataRequestDto = sync.Pool{
	New: func() any {
		return new(SyncHotelBookingDataRequestDto)
	},
}

// GetSyncHotelBookingDataRequestDto() 从对象池中获取SyncHotelBookingDataRequestDto
func GetSyncHotelBookingDataRequestDto() *SyncHotelBookingDataRequestDto {
	return poolSyncHotelBookingDataRequestDto.Get().(*SyncHotelBookingDataRequestDto)
}

// ReleaseSyncHotelBookingDataRequestDto 释放SyncHotelBookingDataRequestDto
func ReleaseSyncHotelBookingDataRequestDto(v *SyncHotelBookingDataRequestDto) {
	v.Segments = v.Segments[:0]
	poolSyncHotelBookingDataRequestDto.Put(v)
}
