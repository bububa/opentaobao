package alitripmerchant

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// HotelDetailInfoDto 结构体
type HotelDetailInfoDto struct {
	// 酒店图片信息
	HotelPictures []HotelPictureDto `json:"hotel_pictures,omitempty" xml:"hotel_pictures>hotel_picture_dto,omitempty"`
	// 娱乐设施图片
	FunFacilitys []string `json:"fun_facilitys,omitempty" xml:"fun_facilitys>string,omitempty"`
	// 酒店设施
	HotelFacilitys []string `json:"hotel_facilitys,omitempty" xml:"hotel_facilitys>string,omitempty"`
	// 酒店服务
	HotelServices []string `json:"hotel_services,omitempty" xml:"hotel_services>string,omitempty"`
	// 房型详情
	RoomDetails []RoomDetailDto `json:"room_details,omitempty" xml:"room_details>room_detail_dto,omitempty"`
	// 酒店政策
	HotelPolicys []string `json:"hotel_policys,omitempty" xml:"hotel_policys>string,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 邮编
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 楼层信息
	Floors string `json:"floors,omitempty" xml:"floors,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 品牌编码
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 房型设施
	Facilitys string `json:"facilitys,omitempty" xml:"facilitys,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 星级
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 酒店中文名称
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 外部酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 最早入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 酒店电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 地区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 最晚离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 国家中文名
	CountryCn string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
	// 评分
	RatingAverage string `json:"rating_average,omitempty" xml:"rating_average,omitempty"`
	// 国家英文名
	CountryEn string `json:"country_en,omitempty" xml:"country_en,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 旗舰店与标准酒店的绑定ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 地区code
	DistrictCode int64 `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 经纬度类型
	PositionType *model.File `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 国别
	Domestic *model.File `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 国家
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省份code
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 0-营业中；-1，筹建中；-2，暂停营业；-3，已停业；-4，失效，-5 ，需电话咨询；默认为0
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
}

var poolHotelDetailInfoDto = sync.Pool{
	New: func() any {
		return new(HotelDetailInfoDto)
	},
}

// GetHotelDetailInfoDto() 从对象池中获取HotelDetailInfoDto
func GetHotelDetailInfoDto() *HotelDetailInfoDto {
	return poolHotelDetailInfoDto.Get().(*HotelDetailInfoDto)
}

// ReleaseHotelDetailInfoDto 释放HotelDetailInfoDto
func ReleaseHotelDetailInfoDto(v *HotelDetailInfoDto) {
	v.HotelPictures = v.HotelPictures[:0]
	v.FunFacilitys = v.FunFacilitys[:0]
	v.HotelFacilitys = v.HotelFacilitys[:0]
	v.HotelServices = v.HotelServices[:0]
	v.RoomDetails = v.RoomDetails[:0]
	v.HotelPolicys = v.HotelPolicys[:0]
	v.OpeningTime = ""
	v.PostalCode = ""
	v.Latitude = ""
	v.Description = ""
	v.Floors = ""
	v.Province = ""
	v.CityCn = ""
	v.Fax = ""
	v.BrandCode = ""
	v.Facilitys = ""
	v.Longitude = ""
	v.BrandName = ""
	v.Address = ""
	v.Star = ""
	v.DecorateTime = ""
	v.NameCn = ""
	v.HotelId = ""
	v.CheckIn = ""
	v.Phone = ""
	v.District = ""
	v.CheckOut = ""
	v.CountryCn = ""
	v.RatingAverage = ""
	v.CountryEn = ""
	v.Rooms = 0
	v.Hid = 0
	v.DistrictCode = 0
	v.PositionType = nil
	v.Shid = 0
	v.CityCode = 0
	v.Domestic = nil
	v.CountryCode = 0
	v.ProvinceCode = 0
	v.Status = nil
	poolHotelDetailInfoDto.Put(v)
}
