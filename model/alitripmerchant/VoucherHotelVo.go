package alitripmerchant

// VoucherHotelVo 结构体
type VoucherHotelVo struct {
	// 房型设施
	HotelFacilityListVOList []FacilityListVo `json:"hotel_facility_list_v_o_list,omitempty" xml:"hotel_facility_list_v_o_list>facility_list_vo,omitempty"`
	// 日期以及价格
	DailyPrices []DailyPrice `json:"daily_prices,omitempty" xml:"daily_prices>daily_price,omitempty"`
	// 租户酒店ID
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店名
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 酒店电话
	HotelPhone string `json:"hotel_phone,omitempty" xml:"hotel_phone,omitempty"`
	// 酒店图片
	HotelPicture string `json:"hotel_picture,omitempty" xml:"hotel_picture,omitempty"`
	// 酒店位置
	HotelAddress string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// 经度
	Lon string `json:"lon,omitempty" xml:"lon,omitempty"`
	// 维度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 入住须知
	CheckInNotice string `json:"check_in_notice,omitempty" xml:"check_in_notice,omitempty"`
	// 取消描述
	CancelDec string `json:"cancel_dec,omitempty" xml:"cancel_dec,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 最早入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 最晚离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 限制时间
	LimitDate string `json:"limit_date,omitempty" xml:"limit_date,omitempty"`
	// 货币类型
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 总房价(不含税)
	TotalRoomPrice string `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 总税价
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 总费用（税费+房价）
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 加价描述
	MarkupDesc string `json:"markup_desc,omitempty" xml:"markup_desc,omitempty"`
	// 外部酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 房间数量
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 取消规则
	CancelRule int64 `json:"cancel_rule,omitempty" xml:"cancel_rule,omitempty"`
	// 入住天数
	Days int64 `json:"days,omitempty" xml:"days,omitempty"`
	// 酒店内房型信息
	Room *Room `json:"room,omitempty" xml:"room,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 房间数量
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// 差价
	Spread int64 `json:"spread,omitempty" xml:"spread,omitempty"`
	// 外币信息
	ForeignCurrency *ForeignCurrencyInfo `json:"foreign_currency,omitempty" xml:"foreign_currency,omitempty"`
	// 是否为外币支付
	ForeignCurrencyPayment bool `json:"foreign_currency_payment,omitempty" xml:"foreign_currency_payment,omitempty"`
}
