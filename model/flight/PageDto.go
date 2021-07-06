package flight

// PageDto 结构体
type PageDto struct {
	// 结果集
	DataList []AlitripAgentFlightSellModifyListT `json:"data_list,omitempty" xml:"data_list>alitrip_agent_flight_sell_modify_list_t,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 123
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
