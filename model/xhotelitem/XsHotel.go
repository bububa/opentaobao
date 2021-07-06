package xhotelitem

// XsHotel 结构体
type XsHotel struct {
	// 国家编码
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 商圈
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 英文名称
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 照片,json格式
	Pics string `json:"pics,omitempty" xml:"pics,omitempty"`
	// 交通距离与设施服务。JSON格式。
	Service string `json:"service,omitempty" xml:"service,omitempty"`
	// 酒店设施。json格式
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 楼层
	Floors string `json:"floors,omitempty" xml:"floors,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 开业年份
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 门店电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 精度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 系统自动生成
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 门店状态：0: 正常;-2:停售；-1：删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 省份编码
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 区域编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 0:国内;1:国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// position_type
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
}
