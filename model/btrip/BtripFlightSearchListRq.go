package btrip

// BtripFlightSearchListRq 结构体
type BtripFlightSearchListRq struct {
	// 忽略店铺列表
	IgnoredShopNames []string `json:"ignored_shop_names,omitempty" xml:"ignored_shop_names>string,omitempty"`
	// 指定店铺列表
	ShopNames []string `json:"shop_names,omitempty" xml:"shop_names>string,omitempty"`
	// 乘客列表
	TravelerList []TravelerInfo `json:"traveler_list,omitempty" xml:"traveler_list>traveler_info,omitempty"`
	// 航司编码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 到达城市三字码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 到达城市名称
	ArrCityName string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// 到达日期YYYY-MM-dd
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 所有舱位(0), 经济舱(1), 商务舱(2), 头等舱(3), 头等舱和商务舱(4), 经济舱和商务舱(5), 经济舱和头等舱(6);
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 出发城市三字码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 出发城市名称
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 出发日期 YYYY-MM-dd
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 乘客数量
	PassengerNum string `json:"passenger_num,omitempty" xml:"passenger_num,omitempty"`
	// 子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 行程类型 0:单程 1:往返
	TripType string `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 可选项,航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 可选项,航班号,中转航班号，（即第二程航班）
	TransferFlightNo string `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	// 可选项，自定义渠道名称（用于创建渠道）
	IsvName string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// 供应商标识
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 中转城市三字码:transferFlightNo不为空则必填
	TransferCityCode string `json:"transfer_city_code,omitempty" xml:"transfer_city_code,omitempty"`
	// 中转航班出发时间:transferFlightNo不为空则必填
	TransferLeaveDate string `json:"transfer_leave_date,omitempty" xml:"transfer_leave_date,omitempty"`
	// 可选项,是否查询多舱价位
	NeedMultiClassRice bool `json:"need_multi_class_rice,omitempty" xml:"need_multi_class_rice,omitempty"`
	// 是否查询多舱价位
	NeedMultiClassPrice bool `json:"need_multi_class_price,omitempty" xml:"need_multi_class_price,omitempty"`
}
