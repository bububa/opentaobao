package btrip

// OpenItineraryInfo 结构体
type OpenItineraryInfo struct {
	// 到达日期
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 发票抬头
	InvoiceName string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	// 成本中心名称
	CostCenterName string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 行程id
	ItineraryId string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// 出发城市编码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 到达城市编码
	ArrCityCode string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// 第三方成本中心id，若不填则商旅成本中心id必填
	ThirdpartCostCenterId string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// 项目代号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 项目名称
	ProjectTitle string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// 第三方发票id，和商旅发票id二者选择其一即可
	ThirdPartInvoiceId string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// 交通方式 0飞机, 1,火车, 2汽车, 3其他
	TrafficType int64 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	// 行程方式：0单程，1往返
	TripWay int64 `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
	// 发票id
	InvoiceId int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 商旅成本中心id，若不填则第三方成本中心id必填
	CostCenterId int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
}
