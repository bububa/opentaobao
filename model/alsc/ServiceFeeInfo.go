package alsc

// ServiceFeeInfo 结构体
type ServiceFeeInfo struct {
	// 服务费实收金额
	ServiceActualFee int64 `json:"service_actual_fee,omitempty" xml:"service_actual_fee,omitempty"`
	// 服务费名称
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 服务费优惠金额，减钱为负值
	ServicePromoFee int64 `json:"service_promo_fee,omitempty" xml:"service_promo_fee,omitempty"`
	// 服务费总金额
	ServiceTotalFee int64 `json:"service_total_fee,omitempty" xml:"service_total_fee,omitempty"`
}
