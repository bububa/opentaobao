package hotel

import (
	"sync"
)

// SHotelVo 结构体
type SHotelVo struct {
	// 标准房型列表
	SroomTypes []SRoomTypeVo `json:"sroom_types,omitempty" xml:"sroom_types>s_room_type_vo,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 品牌code
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 最早入住时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 最晚离店时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 酒店描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// h5详情页url
	H5DetailUrl string `json:"h5_detail_url,omitempty" xml:"h5_detail_url,omitempty"`
	// 默认PC详情页url
	HotelDetailUrl string `json:"hotel_detail_url,omitempty" xml:"hotel_detail_url,omitempty"`
	// 酒店设施文案
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 维度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 酒店中文名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 重复信息
	PetInfo string `json:"pet_info,omitempty" xml:"pet_info,omitempty"`
	// 图片信息
	PicUrls string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
	// 评分
	RateScore string `json:"rate_score,omitempty" xml:"rate_score,omitempty"`
	// 档次，0--无档次，2--二星及经济，3--舒适，4--高档，5--豪华
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 酒店业务分类
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 当前酒店所在城市
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 与某个POI的距离
	Distance int64 `json:"distance,omitempty" xml:"distance,omitempty"`
	// 区县
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 档次
	Level *NameValuePair `json:"level,omitempty" xml:"level,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 省份
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 评论数
	RateNumber int64 `json:"rate_number,omitempty" xml:"rate_number,omitempty"`
	// 销量
	Sell int64 `json:"sell,omitempty" xml:"sell,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 服务设施集合类
	ShotelPropertiesVo *ShotelPropertiesSetVo `json:"shotel_properties_vo,omitempty" xml:"shotel_properties_vo,omitempty"`
	// 上下架状态，0--下架，其他状态-下架
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 客房数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 是否是民宿
	BnbHotel bool `json:"bnb_hotel,omitempty" xml:"bnb_hotel,omitempty"`
}

var poolSHotelVo = sync.Pool{
	New: func() any {
		return new(SHotelVo)
	},
}

// GetSHotelVo() 从对象池中获取SHotelVo
func GetSHotelVo() *SHotelVo {
	return poolSHotelVo.Get().(*SHotelVo)
}

// ReleaseSHotelVo 释放SHotelVo
func ReleaseSHotelVo(v *SHotelVo) {
	v.SroomTypes = v.SroomTypes[:0]
	v.Address = ""
	v.Brand = ""
	v.CheckInTime = ""
	v.CheckOutTime = ""
	v.DecorateTime = ""
	v.Desc = ""
	v.GmtModified = ""
	v.H5DetailUrl = ""
	v.HotelDetailUrl = ""
	v.HotelFacilities = ""
	v.Lat = ""
	v.Lng = ""
	v.Name = ""
	v.OpeningTime = ""
	v.PetInfo = ""
	v.PicUrls = ""
	v.RateScore = ""
	v.Star = ""
	v.Tel = ""
	v.Type = ""
	v.City = 0
	v.Distance = 0
	v.District = 0
	v.Level = nil
	v.Price = 0
	v.Province = 0
	v.RateNumber = 0
	v.Sell = 0
	v.Shid = 0
	v.ShotelPropertiesVo = nil
	v.Status = 0
	v.Rooms = 0
	v.BnbHotel = false
	poolSHotelVo.Put(v)
}
