package ascp

// PromiseDesRule 结构体
type PromiseDesRule struct {
	// 线路发货快递公司编码列表
	PlatformDeliveryCodes []string `json:"platform_delivery_codes,omitempty" xml:"platform_delivery_codes>string,omitempty"`
	// 波次能力
	Waveability []string `json:"waveability,omitempty" xml:"waveability>string,omitempty"`
	// 收货地：省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货地：城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货地：区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地：街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 外部波次id
	OuterWaveId string `json:"outer_wave_id,omitempty" xml:"outer_wave_id,omitempty"`
	// 波次接单截单时间（HH:mm）
	ReceiveCutTime string `json:"receive_cut_time,omitempty" xml:"receive_cut_time,omitempty"`
	// 仓计划出库时间（HH:mm）
	PlanDeliveryTime string `json:"plan_delivery_time,omitempty" xml:"plan_delivery_time,omitempty"`
	// 计划揽收时间（HH:mm）
	PlanCollectionTime string `json:"plan_collection_time,omitempty" xml:"plan_collection_time,omitempty"`
	// 计划入首分拨时间（HH:mm）
	FirstAllocationTime string `json:"first_allocation_time,omitempty" xml:"first_allocation_time,omitempty"`
	// 承诺到达时效类型（不区分快递公司）： 0=不表达 1=承诺当日达 2=承诺次日达 3=承诺隔日达 业务身份为时效代运营时，必填
	TimingType int64 `json:"timing_type,omitempty" xml:"timing_type,omitempty"`
	// 配揽收-签收时长（小于等于）（h）
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 仓接单-签收时长（小于等于）（h）
	ExpressTime int64 `json:"express_time,omitempty" xml:"express_time,omitempty"`
}
