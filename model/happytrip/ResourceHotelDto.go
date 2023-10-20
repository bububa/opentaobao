package happytrip

import (
	"sync"
)

// ResourceHotelDto 结构体
type ResourceHotelDto struct {
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离开时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 城市中文名称
	CityChineseName string `json:"city_chinese_name,omitempty" xml:"city_chinese_name,omitempty"`
	// 城市Code
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 酒店地址
	HotelAddress string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// 酒店中文名称
	HotelChineseName string `json:"hotel_chinese_name,omitempty" xml:"hotel_chinese_name,omitempty"`
	// 酒店英文名称
	HotelEnglishName string `json:"hotel_english_name,omitempty" xml:"hotel_english_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 酒店电话
	HotelTel string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// 酒店GPS位置纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店GPS位置经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 外部酒店id
	OutHotelId string `json:"out_hotel_id,omitempty" xml:"out_hotel_id,omitempty"`
	// 省或州
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 外系统订单号
	ResourceId string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 房间数量
	RoomCount string `json:"room_count,omitempty" xml:"room_count,omitempty"`
	// 房型id
	RoomTypeId string `json:"room_type_id,omitempty" xml:"room_type_id,omitempty"`
	// 房名名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// 酒店ID
	HotelId int64 `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 入住几晚
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 所属订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 资源表主键
	ResourceMainId int64 `json:"resource_main_id,omitempty" xml:"resource_main_id,omitempty"`
}

var poolResourceHotelDto = sync.Pool{
	New: func() any {
		return new(ResourceHotelDto)
	},
}

// GetResourceHotelDto() 从对象池中获取ResourceHotelDto
func GetResourceHotelDto() *ResourceHotelDto {
	return poolResourceHotelDto.Get().(*ResourceHotelDto)
}

// ReleaseResourceHotelDto 释放ResourceHotelDto
func ReleaseResourceHotelDto(v *ResourceHotelDto) {
	v.CheckIn = ""
	v.CheckOut = ""
	v.CityChineseName = ""
	v.CityCode = ""
	v.CityName = ""
	v.Country = ""
	v.District = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.HotelAddress = ""
	v.HotelChineseName = ""
	v.HotelEnglishName = ""
	v.HotelName = ""
	v.HotelTel = ""
	v.Latitude = ""
	v.Longitude = ""
	v.OutHotelId = ""
	v.Province = ""
	v.ResourceId = ""
	v.RoomCount = ""
	v.RoomTypeId = ""
	v.RoomTypeName = ""
	v.HotelId = 0
	v.Id = 0
	v.Nights = 0
	v.OrderId = 0
	v.ResourceMainId = 0
	poolResourceHotelDto.Put(v)
}
