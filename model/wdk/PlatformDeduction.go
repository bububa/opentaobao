package wdk

// PlatformDeduction 结构体
type PlatformDeduction struct {
	// 技术服务费
	TechnicalServiceFee int64 `json:"technical_service_fee,omitempty" xml:"technical_service_fee,omitempty"`
	// 支付服务费
	PayServiceFee int64 `json:"pay_service_fee,omitempty" xml:"pay_service_fee,omitempty"`
	// 基础物流费
	BaseLogisticsFee int64 `json:"base_logistics_fee,omitempty" xml:"base_logistics_fee,omitempty"`
	// 代运营服务费
	ThirdpartnarFee int64 `json:"thirdpartnar_fee,omitempty" xml:"thirdpartnar_fee,omitempty"`
}
