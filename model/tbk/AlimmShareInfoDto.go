package tbk

// AlimmShareInfoDto 结构体
type AlimmShareInfoDto struct {
	// 技术服务费比率
	Alimmtechservicerate float64 `json:"alimm_tech_service_rate,omitempty" xml:"alimm_tech_service_rate,omitempty"`
	// 预估技术服务费
	Alimmtechserviceprefee float64 `json:"alimm_tech_service_pre_fee,omitempty" xml:"alimm_tech_service_pre_fee,omitempty"`
	// 结算技术服务费
	Alimmtechservicefee float64 `json:"alimm_tech_service_fee,omitempty" xml:"alimm_tech_service_fee,omitempty"`
	// 渠道专项服务费比率
	Alimmagentservicerate float64 `json:"alimm_agent_service_rate,omitempty" xml:"alimm_agent_service_rate,omitempty"`
	// 预估渠道专项服务费
	Alimmagentserviceprefee float64 `json:"alimm_agent_service_pre_fee,omitempty" xml:"alimm_agent_service_pre_fee,omitempty"`
	// 结算渠道专项服务费
	Alimmagentservicefee float64 `json:"alimm_agent_service_fee,omitempty" xml:"alimm_agent_service_fee,omitempty"`
}
