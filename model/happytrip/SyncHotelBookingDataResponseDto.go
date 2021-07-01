package happytrip

// SyncHotelBookingDataResponseDto 结构体
type SyncHotelBookingDataResponseDto struct {
	// 消息发送结果
	Segments []HotelMessageSendSegment `json:"segments,omitempty" xml:"segments>hotel_message_send_segment,omitempty"`
}
