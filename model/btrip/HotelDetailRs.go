package btrip

// HotelDetailRs 
type HotelDetailRs struct {
    // 入住时间
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    // 离店时间
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    // 跟踪ID，可忽略
    EagleTraceId   string `json:"eagle_trace_id,omitempty" xml:"eagle_trace_id,omitempty"`
    // 房型列表
    Rooms   []HotelDetailRoomDTO `json:"rooms,omitempty" xml:"rooms>hotel_detail_room_dto,omitempty"`
    // 跟踪ID。可忽略
    SearchId   string `json:"search_id,omitempty" xml:"search_id,omitempty"`
    // 酒店标准ID
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
