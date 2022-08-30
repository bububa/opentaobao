package axindata

import (
	"github.com/bububa/opentaobao/model"
)

// StdHotelVO 结构体
type StdHotelVO struct {
	// 标准房型信息
	StdRoomTypeList []StdRoomType `json:"std_room_type_list,omitempty" xml:"std_room_type_list>std_room_type,omitempty"`
	// 标准酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 服务
	Services string `json:"services,omitempty" xml:"services,omitempty"`
	// 经度
	Longtitude string `json:"longtitude,omitempty" xml:"longtitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 评分
	RateScore string `json:"rate_score,omitempty" xml:"rate_score,omitempty"`
	// 酒店电话
	HotelTel string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// 酒店特色
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 酒店档次
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 酒店设施
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 酒店品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 宠物信息
	PetInfo string `json:"pet_info,omitempty" xml:"pet_info,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 评论数
	RateNumber int64 `json:"rate_number,omitempty" xml:"rate_number,omitempty"`
	// 城市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 酒店类型（0-标准酒店 1-标准客栈）
	HotelType *model.File `json:"hotel_type,omitempty" xml:"hotel_type,omitempty"`
	// 状态0-正常
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
}
