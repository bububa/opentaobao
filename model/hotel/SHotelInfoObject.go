package hotel

// SHotelInfoObject 结构体
type SHotelInfoObject struct {
	// 房型信息
	Rooms []SRoomType `json:"rooms,omitempty" xml:"rooms>s_room_type,omitempty"`
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 酒店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 纬度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 酒店星级，1-5星，0是客栈
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 酒店开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 酒店装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 电话，包括三种类型：<br/>1.固定电话，例如：0086-010-85322688<br/>2.移动电话，例如：13869696363<br/>3.400或800电话，例如：0086-4006123928
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 评分
	RateScore string `json:"rate_score,omitempty" xml:"rate_score,omitempty"`
	// 酒店描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 图片url，多张图片使用","隔开
	PicUrls string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
	// 酒店设施
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 服务设施
	Services string `json:"services,omitempty" xml:"services,omitempty"`
	// 酒店品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// H5的detail页面的URL
	H5DetailUrl string `json:"h5_detail_url,omitempty" xml:"h5_detail_url,omitempty"`
	// 酒店detail页面的url
	PcDetailUrl string `json:"pc_detail_url,omitempty" xml:"pc_detail_url,omitempty"`
	// 入住时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 离店时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 宠物信息
	PetInfo string `json:"pet_info,omitempty" xml:"pet_info,omitempty"`
	// 外宾描述
	ForeignDesc string `json:"foreign_desc,omitempty" xml:"foreign_desc,omitempty"`
	// 标准酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 省的code
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 市的code
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 地区的值
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 评论数
	RateNumber int64 `json:"rate_number,omitempty" xml:"rate_number,omitempty"`
	// 酒店状态,0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 外宾类型
	ForeignType int64 `json:"foreign_type,omitempty" xml:"foreign_type,omitempty"`
	// 是否为民宿类型
	BnbHotel bool `json:"bnb_hotel,omitempty" xml:"bnb_hotel,omitempty"`
}
