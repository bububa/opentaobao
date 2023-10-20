package tbk

// AlimmShareInfoDto 结构体
type AlimmShareInfoDto struct {
	// 技术服务费比率
	AlimmTechServiceRate float64 `json:"alimm_tech_service_rate,omitempty" xml:"alimm_tech_service_rate,omitempty"`
	// 预估技术服务费
	AlimmTechServicePreFee float64 `json:"alimm_tech_service_pre_fee,omitempty" xml:"alimm_tech_service_pre_fee,omitempty"`
	// 结算技术服务费
	AlimmTechServiceFee float64 `json:"alimm_tech_service_fee,omitempty" xml:"alimm_tech_service_fee,omitempty"`
	// 渠道专项服务费比率
	AlimmAgentServiceRate float64 `json:"alimm_agent_service_rate,omitempty" xml:"alimm_agent_service_rate,omitempty"`
	// 预估渠道专项服务费
	AlimmAgentServicePreFee float64 `json:"alimm_agent_service_pre_fee,omitempty" xml:"alimm_agent_service_pre_fee,omitempty"`
	// 结算渠道专项服务费
	AlimmAgentServiceFee float64 `json:"alimm_agent_service_fee,omitempty" xml:"alimm_agent_service_fee,omitempty"`
}
