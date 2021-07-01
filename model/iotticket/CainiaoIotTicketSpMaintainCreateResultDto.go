package iotticket

// CainiaoIotTicketSpMaintainCreateResultDto 结构体
type CainiaoIotTicketSpMaintainCreateResultDto struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
