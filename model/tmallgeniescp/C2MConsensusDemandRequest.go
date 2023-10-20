package tmallgeniescp

// C2mconsensusDemandRequest 结构体
type C2mconsensusDemandRequest struct {
	// 对象
	C2MConsensusDemandParamDTOList []C2mconsensusDemandParamDto `json:"c2_m_consensus_demand_param_d_t_o_list,omitempty" xml:"c2_m_consensus_demand_param_d_t_o_list>c2mconsensus_demand_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
