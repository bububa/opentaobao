package flight

// AlitripAgentFlightSellRefundRefuseResultDto 
/* model for simplify = false
type AlitripAgentFlightSellRefundRefuseResultDto struct {

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// AlitripAgentFlightSellRefundRefuseResultDto 
type AlitripAgentFlightSellRefundRefuseResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
