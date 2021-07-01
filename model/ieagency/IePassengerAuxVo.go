package ieagency

// IePassengerAuxVo 结构体
type IePassengerAuxVo struct {
	// 乘机人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 乘机人购买辅营产品的航段
	FlightVo *IeBookFlightVo `json:"flight_vo,omitempty" xml:"flight_vo,omitempty"`
	// 辅营产品规格信息
	AuxProductVo *IeAuxProductVo `json:"aux_product_vo,omitempty" xml:"aux_product_vo,omitempty"`
	// 订购数量，默认1
	BookNum int64 `json:"book_num,omitempty" xml:"book_num,omitempty"`
}
