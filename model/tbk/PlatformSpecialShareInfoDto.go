package tbk

// PlatformSpecialShareInfoDto 结构体
type PlatformSpecialShareInfoDto struct {
	// 内容专项服务费比率
	ContentTechServiceRate float64 `json:"content_tech_service_rate,omitempty" xml:"content_tech_service_rate,omitempty"`
	// 预估内容专项服务费
	ContentTechServicePreFee float64 `json:"content_tech_service_pre_fee,omitempty" xml:"content_tech_service_pre_fee,omitempty"`
	// 结算内容专项服务费
	ContentTechServiceFee float64 `json:"content_tech_service_fee,omitempty" xml:"content_tech_service_fee,omitempty"`
	// 流量专项服务费比率（默认无，限定开放）
	TrafficTechServiceRate float64 `json:"traffic_tech_service_rate,omitempty" xml:"traffic_tech_service_rate,omitempty"`
	// 预估流量专项服务费（默认无，限定开放）
	TrafficTechServicePreFee float64 `json:"traffic_tech_service_pre_fee,omitempty" xml:"traffic_tech_service_pre_fee,omitempty"`
	// 结算流量专项服务费（默认无，限定开放）
	TrafficTechServiceFee float64 `json:"traffic_tech_service_fee,omitempty" xml:"traffic_tech_service_fee,omitempty"`
}
