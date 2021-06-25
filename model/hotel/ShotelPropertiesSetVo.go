package hotel

// ShotelPropertiesSetVo 
type ShotelPropertiesSetVo struct {

    // 预订须知
    ShotelBookingNotics   []ShotelPropertiesVo `json:"shotel_booking_notics,omitempty"`

    // 娱乐设施
    ShotelFunFacilities   []ShotelPropertiesVo `json:"shotel_fun_facilities,omitempty"`

    // 酒店设施
    ShotelHotelFacilities   []ShotelPropertiesVo `json:"shotel_hotel_facilities,omitempty"`

    // 酒店服务
    ShotelHotelServices   []ShotelPropertiesVo `json:"shotel_hotel_services,omitempty"`

    // 酒店普通图
    ShotelNomalPictures   []ShotelPropertiesVo `json:"shotel_nomal_pictures,omitempty"`

    // 酒店维度的房间设施
    ShotelRoomFacilities   []ShotelPropertiesVo `json:"shotel_room_facilities,omitempty"`

}
