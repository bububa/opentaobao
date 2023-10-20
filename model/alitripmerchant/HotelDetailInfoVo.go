package alitripmerchant

import (
	"sync"
)

// HotelDetailInfoVo 结构体
type HotelDetailInfoVo struct {
	// 设施组
	FacilityGroupList []FacilityListVo `json:"facility_group_list,omitempty" xml:"facility_group_list>facility_list_vo,omitempty"`
	// 价格分组概要
	PriceGroupSummaryList []PriceGroupSummaryVo `json:"price_group_summary_list,omitempty" xml:"price_group_summary_list>price_group_summary_vo,omitempty"`
	// 酒店图片
	HotelPictures []HotelPictureDto `json:"hotel_pictures,omitempty" xml:"hotel_pictures>hotel_picture_dto,omitempty"`
	// 房型详情
	RoomDetails []RoomDetailVo `json:"room_details,omitempty" xml:"room_details>room_detail_vo,omitempty"`
	// 酒店政策集合
	HotelPolicyList []FacilityVo `json:"hotel_policy_list,omitempty" xml:"hotel_policy_list>facility_vo,omitempty"`
	// 酒店中文名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店地址信息
	AddressInfo *AddressVo `json:"address_info,omitempty" xml:"address_info,omitempty"`
	// 酒店信息
	HotelInfo *HotelVo `json:"hotel_info,omitempty" xml:"hotel_info,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 卫生健康标识
	HealthNotification *FacilityVo `json:"health_notification,omitempty" xml:"health_notification,omitempty"`
	// 飞猪旗舰店的ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
}

var poolHotelDetailInfoVo = sync.Pool{
	New: func() any {
		return new(HotelDetailInfoVo)
	},
}

// GetHotelDetailInfoVo() 从对象池中获取HotelDetailInfoVo
func GetHotelDetailInfoVo() *HotelDetailInfoVo {
	return poolHotelDetailInfoVo.Get().(*HotelDetailInfoVo)
}

// ReleaseHotelDetailInfoVo 释放HotelDetailInfoVo
func ReleaseHotelDetailInfoVo(v *HotelDetailInfoVo) {
	v.FacilityGroupList = v.FacilityGroupList[:0]
	v.PriceGroupSummaryList = v.PriceGroupSummaryList[:0]
	v.HotelPictures = v.HotelPictures[:0]
	v.RoomDetails = v.RoomDetails[:0]
	v.HotelPolicyList = v.HotelPolicyList[:0]
	v.NameCn = ""
	v.HotelId = ""
	v.AddressInfo = nil
	v.HotelInfo = nil
	v.Shid = 0
	v.HealthNotification = nil
	v.Hid = 0
	poolHotelDetailInfoVo.Put(v)
}
