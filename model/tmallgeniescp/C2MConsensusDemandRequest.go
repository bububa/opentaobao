package tmallgeniescp

// C2MConsensusDemandRequest 结构体
type C2MConsensusDemandRequest struct {
	// 对象
	C2MConsensusDemandParamDTOList []C2MConsensusDemandParamDto `json:"c2_m_consensus_demand_param_d_t_o_list,omitempty" xml:"c2_m_consensus_demand_param_d_t_o_list>c2m_consensus_demand_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
