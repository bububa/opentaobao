package flight

// AlitripAgentFlightSellTicketingDetailResultDto 
/* model for simplify = false
type AlitripAgentFlightSellTicketingDetailResultDto struct {

    // 出参对象
    
    Data  *struct {
        AlitripAgentFlightSellTicketingDetailData  *AlitripAgentFlightSellTicketingDetailData `json:"alitrip_agent_flight_sell_ticketing_detail_data,omitempty"`
    } `json:"data,omitempty"`
    

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// AlitripAgentFlightSellTicketingDetailResultDto 
type AlitripAgentFlightSellTicketingDetailResultDto struct {

    // 出参对象
    Data   *AlitripAgentFlightSellTicketingDetailData `json:"data,omitempty"`

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
