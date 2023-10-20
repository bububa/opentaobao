package happytrip

import (
	"sync"
)

// SyncHotelBookingDataResponseDto 结构体
type SyncHotelBookingDataResponseDto struct {
	// 消息发送结果
	Segments []HotelMessageSendSegment `json:"segments,omitempty" xml:"segments>hotel_message_send_segment,omitempty"`
}

var poolSyncHotelBookingDataResponseDto = sync.Pool{
	New: func() any {
		return new(SyncHotelBookingDataResponseDto)
	},
}

// GetSyncHotelBookingDataResponseDto() 从对象池中获取SyncHotelBookingDataResponseDto
func GetSyncHotelBookingDataResponseDto() *SyncHotelBookingDataResponseDto {
	return poolSyncHotelBookingDataResponseDto.Get().(*SyncHotelBookingDataResponseDto)
}

// ReleaseSyncHotelBookingDataResponseDto 释放SyncHotelBookingDataResponseDto
func ReleaseSyncHotelBookingDataResponseDto(v *SyncHotelBookingDataResponseDto) {
	v.Segments = v.Segments[:0]
	poolSyncHotelBookingDataResponseDto.Put(v)
}
