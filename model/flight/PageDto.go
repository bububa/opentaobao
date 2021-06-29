package flight

// PageDto 
type PageDto struct {
    // 123
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 结果集
    DataList   []AlitripAgentFlightSellModifyListT `json:"data_list,omitempty" xml:"data_list>alitrip_agent_flight_sell_modify_list_t,omitempty"`
    // 成功标识
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
