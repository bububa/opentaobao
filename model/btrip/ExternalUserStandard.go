package btrip

// ExternalUserStandard 结构体
type ExternalUserStandard struct {
	// 酒店城市费用列表
	HotelCitys []HotelCityFee `json:"hotel_citys,omitempty" xml:"hotel_citys>hotel_city_fee,omitempty"`
	// 国内机票舱等，多值逗号分隔。F：头等舱，C：商务舱，Y：经济舱
	FlightCabins string `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	// 火车票坐席，多值逗号分隔。0：硬座，1：硬卧，2：软座，3：软卧，4：高级软卧，5：商务座，6：一等座，7：二等座，8：动卧
	TrainSeats string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	// 国际机票舱等，多值逗号分隔。F：头等舱，C：商务舱，Y：经济舱，P：超级经济舱
	InternationalFlightCabins string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	// 经济舱折扣。1到10的整数
	EconomyDiscount int64 `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	// 商务舱折扣。1到10的整数
	BusinessDiscount int64 `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	// 头等舱折扣。1到10的整数
	FirstDiscount int64 `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	// 限制差标类型。0-不限差标，1-限制差标。注意：同一审批单的所有出行人只能都限制差标/都不限制差标，否则会调用失败
	ReserveType int64 `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
}
