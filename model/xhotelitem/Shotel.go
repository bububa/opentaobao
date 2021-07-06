package xhotelitem

// SHotel 结构体
type SHotel struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// used_name
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
	// 国家编码
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店级别
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// business
	Business string `json:"business,omitempty" xml:"business,omitempty"`
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
	// 酒店介绍
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 交通距离与设施服务。JSON格式。
	Service string `json:"service,omitempty" xml:"service,omitempty"`
	// 酒店图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 扩展信息的JSON
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// brand
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 酒店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 酒店设施
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 房间设施
	RoomFacilities string `json:"room_facilities,omitempty" xml:"room_facilities,omitempty"`
	// 地区标签
	CityTag string `json:"city_tag,omitempty" xml:"city_tag,omitempty"`
	// 邮编
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 杂费
	Fee string `json:"fee,omitempty" xml:"fee,omitempty"`
	// pics
	Pics string `json:"pics,omitempty" xml:"pics,omitempty"`
	// area
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// facility
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// floor
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// bed
	Bed string `json:"bed,omitempty" xml:"bed,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// networkService
	NetworkService string `json:"network_service,omitempty" xml:"network_service,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// windowType
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
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
	// 状态:0：待系统匹配1：已系统匹配，匹配成功，待卖家确认2：已系统匹配，匹配失败，待人工匹配3：已人工匹配，匹配成功，待卖家确认4：已人工匹配，匹配失败5：卖家已确认，确认&ldquo;YES&rdquo;6：卖家已确认，确认&ldquo;NO&rdquo;7：停售
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// srid
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// maxOccupancy
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// addBed
	AddBed int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
	// 0：酒店 1：客栈
	IsKezhan int64 `json:"is_kezhan,omitempty" xml:"is_kezhan,omitempty"`
}
