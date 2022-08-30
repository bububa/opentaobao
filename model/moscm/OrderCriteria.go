package moscm

// OrderCriteria 结构体
type OrderCriteria struct {
	// 订单号
	OrderNumbers []string `json:"order_numbers,omitempty" xml:"order_numbers>string,omitempty"`
	// 未支付(“UNPAID”),已支付("PAID"),部分发货("PARTDISTRIBUTION"),全部发货("ALLDISTRIBUTION"),取消("CANCEL")
	Status []string `json:"status,omitempty" xml:"status>string,omitempty"`
	// 银泰专柜Id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 订单创建时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 订单创建时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 供应商专柜Id
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 订单修改时间范围，开始时间
	ModifyEndDate string `json:"modify_end_date,omitempty" xml:"modify_end_date,omitempty"`
	// 订单修改时间范围，结束时间
	ModifyStartDate string `json:"modify_start_date,omitempty" xml:"modify_start_date,omitempty"`
}
