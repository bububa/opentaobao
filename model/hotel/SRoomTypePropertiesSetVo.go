package hotel

// SRoomTypePropertiesSetVo 
type SRoomTypePropertiesSetVo struct {
    // 标准房型普通图集合
    SroomTypeNomalPictures   []ShotelPropertiesVo `json:"sroom_type_nomal_pictures,omitempty" xml:"sroom_type_nomal_pictures>shotel_properties_vo,omitempty"`
    // 房间设施集合
    SroomTypeRoomFacilities   []ShotelPropertiesVo `json:"sroom_type_room_facilities,omitempty" xml:"sroom_type_room_facilities>shotel_properties_vo,omitempty"`
}
