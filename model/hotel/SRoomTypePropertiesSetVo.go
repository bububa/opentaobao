package hotel

// SRoomTypePropertiesSetVo 
type SRoomTypePropertiesSetVo struct {

    // 标准房型普通图集合
    SroomTypeNomalPictures   []ShotelPropertiesVo `json:"sroom_type_nomal_pictures,omitempty"`

    // 房间设施集合
    SroomTypeRoomFacilities   []ShotelPropertiesVo `json:"sroom_type_room_facilities,omitempty"`

}
