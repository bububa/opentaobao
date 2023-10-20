package hotel

import (
	"sync"
)

// ShotelPropertiesSetVo 结构体
type ShotelPropertiesSetVo struct {
	// 预订须知
	ShotelBookingNotics []ShotelPropertiesVo `json:"shotel_booking_notics,omitempty" xml:"shotel_booking_notics>shotel_properties_vo,omitempty"`
	// 娱乐设施
	ShotelFunFacilities []ShotelPropertiesVo `json:"shotel_fun_facilities,omitempty" xml:"shotel_fun_facilities>shotel_properties_vo,omitempty"`
	// 酒店设施
	ShotelHotelFacilities []ShotelPropertiesVo `json:"shotel_hotel_facilities,omitempty" xml:"shotel_hotel_facilities>shotel_properties_vo,omitempty"`
	// 酒店服务
	ShotelHotelServices []ShotelPropertiesVo `json:"shotel_hotel_services,omitempty" xml:"shotel_hotel_services>shotel_properties_vo,omitempty"`
	// 酒店普通图
	ShotelNomalPictures []ShotelPropertiesVo `json:"shotel_nomal_pictures,omitempty" xml:"shotel_nomal_pictures>shotel_properties_vo,omitempty"`
	// 酒店维度的房间设施
	ShotelRoomFacilities []ShotelPropertiesVo `json:"shotel_room_facilities,omitempty" xml:"shotel_room_facilities>shotel_properties_vo,omitempty"`
}

var poolShotelPropertiesSetVo = sync.Pool{
	New: func() any {
		return new(ShotelPropertiesSetVo)
	},
}

// GetShotelPropertiesSetVo() 从对象池中获取ShotelPropertiesSetVo
func GetShotelPropertiesSetVo() *ShotelPropertiesSetVo {
	return poolShotelPropertiesSetVo.Get().(*ShotelPropertiesSetVo)
}

// ReleaseShotelPropertiesSetVo 释放ShotelPropertiesSetVo
func ReleaseShotelPropertiesSetVo(v *ShotelPropertiesSetVo) {
	v.ShotelBookingNotics = v.ShotelBookingNotics[:0]
	v.ShotelFunFacilities = v.ShotelFunFacilities[:0]
	v.ShotelHotelFacilities = v.ShotelHotelFacilities[:0]
	v.ShotelHotelServices = v.ShotelHotelServices[:0]
	v.ShotelNomalPictures = v.ShotelNomalPictures[:0]
	v.ShotelRoomFacilities = v.ShotelRoomFacilities[:0]
	poolShotelPropertiesSetVo.Put(v)
}
