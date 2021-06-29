package hotelhstdf

// AlitripHotelHstdfShotelRoomtypeMappingsListResults 
type AlitripHotelHstdfShotelRoomtypeMappingsListResults struct {

    // 是否匹配
    
    Matched   bool `json:"matched,omitempty" xml:"matched,omitempty"`
    

    // 卖家房型
    
    RoomType   *RoomTypePo `json:"room_type,omitempty" xml:"room_type,omitempty"`
    

    // 标准房型
    
    StdRoomType   *SroomTypePo `json:"std_room_type,omitempty" xml:"std_room_type,omitempty"`
    

}
