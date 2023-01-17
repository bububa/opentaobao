package btrip

// BtripFlightModifyFlightInfoRq 结构体
type BtripFlightModifyFlightInfoRq struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 改签航班号
	ModifyFlightNo string `json:"modify_flight_no,omitempty" xml:"modify_flight_no,omitempty"`
	// 改签日期
	ModifyDepartDate string `json:"modify_depart_date,omitempty" xml:"modify_depart_date,omitempty"`
	// 出发城市三字码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 到达城市三字码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 分销子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 供应商标识
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	//  航班搜索出来的session
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 改签乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 是否是自愿改签
	IsVoluntary int64 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
}
