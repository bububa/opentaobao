package alscmerchant

import (
	"sync"
)

// AlibabaAlscDaodianTicketReserveResult 结构体
type AlibabaAlscDaodianTicketReserveResult struct {
	// 错误码，success=false时有效
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案，success=false时有效
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 处理结果是否成功，true为成功，false为失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试，success=false时有效
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolAlibabaAlscDaodianTicketReserveResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlscDaodianTicketReserveResult)
	},
}

// GetAlibabaAlscDaodianTicketReserveResult() 从对象池中获取AlibabaAlscDaodianTicketReserveResult
func GetAlibabaAlscDaodianTicketReserveResult() *AlibabaAlscDaodianTicketReserveResult {
	return poolAlibabaAlscDaodianTicketReserveResult.Get().(*AlibabaAlscDaodianTicketReserveResult)
}

// ReleaseAlibabaAlscDaodianTicketReserveResult 释放AlibabaAlscDaodianTicketReserveResult
func ReleaseAlibabaAlscDaodianTicketReserveResult(v *AlibabaAlscDaodianTicketReserveResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Retry = false
	poolAlibabaAlscDaodianTicketReserveResult.Put(v)
}
