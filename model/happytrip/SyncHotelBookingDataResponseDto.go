package happytrip

// SyncHotelBookingDataResponseDto 
type SyncHotelBookingDataResponseDto struct {

    // 消息发送结果
    
    Segments   []HotelMessageSendSegment `json:"segments,omitempty" xml:"segments,omitempty"`
    

}
