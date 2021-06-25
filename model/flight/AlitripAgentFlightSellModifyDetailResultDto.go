package flight

// AlitripAgentFlightSellModifyDetailResultDto 
type AlitripAgentFlightSellModifyDetailResultDto struct {

    // 结果对象
    Data   *ModifyDetailDto `json:"data,omitempty"`

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
