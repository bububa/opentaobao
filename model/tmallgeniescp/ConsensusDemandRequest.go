package tmallgeniescp

// ConsensusDemandRequest 
type ConsensusDemandRequest struct {
    // 入参对象列表
    ConsensusDemandParamDTOList   []ConsensusDemandParamDto `json:"consensus_demand_param_d_t_o_list,omitempty" xml:"consensus_demand_param_d_t_o_list>consensus_demand_param_dto,omitempty"`
    // 扩展参数
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
