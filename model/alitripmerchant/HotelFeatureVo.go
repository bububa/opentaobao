package alitripmerchant

// HotelFeatureVo 结构体
type HotelFeatureVo struct {
	// 入店时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 楼层
	Floors string `json:"floors,omitempty" xml:"floors,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 邮编
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 酒店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
}
