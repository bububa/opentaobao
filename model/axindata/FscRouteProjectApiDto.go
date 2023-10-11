package axindata

// FscRouteProjectApiDto 结构体
type FscRouteProjectApiDto struct {
	// 价格体系
	PriceList []FscRouteProjectPriceApiDto `json:"price_list,omitempty" xml:"price_list>fsc_route_project_price_api_dto,omitempty"`
	// 出团日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 团期计划编号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 截止报名日期
	EndSignDate string `json:"end_sign_date,omitempty" xml:"end_sign_date,omitempty"`
	// 订金单价(单位分)
	BookingUnitPrice int64 `json:"booking_unit_price,omitempty" xml:"booking_unit_price,omitempty"`
	// 剩余库存数量
	InvCount int64 `json:"inv_count,omitempty" xml:"inv_count,omitempty"`
	// 已占库存数量
	OccupyCount int64 `json:"occupy_count,omitempty" xml:"occupy_count,omitempty"`
	// 佣金设置
	FscSaleCommission *FscSaleCommissionApiDto `json:"fsc_sale_commission,omitempty" xml:"fsc_sale_commission,omitempty"`
	// 去程航班
	DepartFlight *FlightInfoApiDto `json:"depart_flight,omitempty" xml:"depart_flight,omitempty"`
	// 返程航班
	ReturnFlight *FlightInfoApiDto `json:"return_flight,omitempty" xml:"return_flight,omitempty"`
}
