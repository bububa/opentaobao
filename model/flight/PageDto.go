package flight

// PageDto 结构体
type PageDto struct {
	// 结果集
	DataList []AlitripAgentCoordinateListT `json:"data_list,omitempty" xml:"data_list>alitrip_agent_coordinate_list_t,omitempty"`
	// 数据集合
	Data []AlitripAgentFlightIntentionListT `json:"data,omitempty" xml:"data>alitrip_agent_flight_intention_list_t,omitempty"`
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 成功标示
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
