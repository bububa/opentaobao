package alitripmerchant

// PriceDetailDto 结构体
type PriceDetailDto struct {
	// 每日价格
	DailyPrices []DailyPrice `json:"daily_prices,omitempty" xml:"daily_prices>daily_price,omitempty"`
	// 房型图片集合
	Pics []string `json:"pics,omitempty" xml:"pics>string,omitempty"`
	// 房型设施
	FacilityGroupList []FacilityListVO `json:"facility_group_list,omitempty" xml:"facility_group_list>facility_list_vo,omitempty"`
	// 总价格
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 总税费
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 总房间价格(不含税)
	TotalRoomPrice string `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 货币
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 夏季限时优惠信息
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 入住时间
	Days int64 `json:"days,omitempty" xml:"days,omitempty"`
	// 房间数量
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// 房间属性
	RoomProperty *DailyPrice `json:"room_property,omitempty" xml:"room_property,omitempty"`
}
