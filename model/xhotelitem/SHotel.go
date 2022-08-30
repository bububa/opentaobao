package xhotelitem

// SHotel 结构体
type SHotel struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// used_name
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
	// 酒店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 国家编码
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 地区标签
	CityTag string `json:"city_tag,omitempty" xml:"city_tag,omitempty"`
	// business
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店级别
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// longitude
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// latitude
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 开业年份
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 装修年份
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 楼层数
	Storeys string `json:"storeys,omitempty" xml:"storeys,omitempty"`
	// 扩展信息的JSON
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 酒店介绍
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 交通距离与设施服务。JSON格式。
	Service string `json:"service,omitempty" xml:"service,omitempty"`
	// 酒店设施
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 房间设施
	RoomFacilities string `json:"room_facilities,omitempty" xml:"room_facilities,omitempty"`
	// 酒店图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 邮编
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// brand
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 杂费
	Fee string `json:"fee,omitempty" xml:"fee,omitempty"`
	// 标准酒店英文名
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 0:国内;1:国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 省份编码
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 区域编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// position_type
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；默认为0
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 0：酒店 1：客栈
	IsKezhan int64 `json:"is_kezhan,omitempty" xml:"is_kezhan,omitempty"`
}
