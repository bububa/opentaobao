package flight

// AlitripAgentFlightSellTicketingDetailResultDto 
type AlitripAgentFlightSellTicketingDetailResultDto struct {

    // 出参对象
    
    Data   *AlitripAgentFlightSellTicketingDetailData `json:"data,omitempty" xml:"data,omitempty"`
    

    // 执行结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
