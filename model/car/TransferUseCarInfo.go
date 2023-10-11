package car

// TransferUseCarInfo 结构体
type TransferUseCarInfo struct {
	// 用车时间
	CarUseTime string `json:"car_use_time,omitempty" xml:"car_use_time,omitempty"`
	// 商家境外客服电话
	AbroadCustomerServicePhone string `json:"abroad_customer_service_phone,omitempty" xml:"abroad_customer_service_phone,omitempty"`
	// 出发地点
	OriginAddress string `json:"origin_address,omitempty" xml:"origin_address,omitempty"`
	// 退改规则
	CancelRule string `json:"cancel_rule,omitempty" xml:"cancel_rule,omitempty"`
	// 车型id
	CarTypeId string `json:"car_type_id,omitempty" xml:"car_type_id,omitempty"`
	// 商家境内客服电话
	DomesticCustomerServicePhone string `json:"domestic_customer_service_phone,omitempty" xml:"domestic_customer_service_phone,omitempty"`
	// 到达地经纬度
	ToLocation string `json:"to_location,omitempty" xml:"to_location,omitempty"`
	// 到达城市
	ArriveCity string `json:"arrive_city,omitempty" xml:"arrive_city,omitempty"`
	// 到达城市三字码
	ArriveCityCode string `json:"arrive_city_code,omitempty" xml:"arrive_city_code,omitempty"`
	// 到达地点
	ArriveAddress string `json:"arrive_address,omitempty" xml:"arrive_address,omitempty"`
	// 出发城市
	OriginCity string `json:"origin_city,omitempty" xml:"origin_city,omitempty"`
	// 出发地经纬度
	FromLocation string `json:"from_location,omitempty" xml:"from_location,omitempty"`
	// 出发城市三字码
	OriginCityCode string `json:"origin_city_code,omitempty" xml:"origin_city_code,omitempty"`
}
