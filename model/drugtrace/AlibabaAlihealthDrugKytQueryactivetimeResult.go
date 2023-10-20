package drugtrace

// AlibabaalihealthdrugkytqueryactivetimeResult 结构体
type AlibabaalihealthdrugkytqueryactivetimeResult struct {
	// 码激活状态DTO
	Models []CodeActiveStatusDto `json:"models,omitempty" xml:"models>code_active_status_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
