package alscmerchant

import (
	"sync"
)

// AlibabaAlscDaodianTicketConsultResult 结构体
type AlibabaAlscDaodianTicketConsultResult struct {
	// 错误码，success=false时有效
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案，success=false时有效
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结构体
	Value *TicketConsultResponse `json:"value,omitempty" xml:"value,omitempty"`
	// 处理结果是否成功，true为成功，false为失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlscDaodianTicketConsultResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlscDaodianTicketConsultResult)
	},
}

// GetAlibabaAlscDaodianTicketConsultResult() 从对象池中获取AlibabaAlscDaodianTicketConsultResult
func GetAlibabaAlscDaodianTicketConsultResult() *AlibabaAlscDaodianTicketConsultResult {
	return poolAlibabaAlscDaodianTicketConsultResult.Get().(*AlibabaAlscDaodianTicketConsultResult)
}

// ReleaseAlibabaAlscDaodianTicketConsultResult 释放AlibabaAlscDaodianTicketConsultResult
func ReleaseAlibabaAlscDaodianTicketConsultResult(v *AlibabaAlscDaodianTicketConsultResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Value = nil
	v.Success = false
	poolAlibabaAlscDaodianTicketConsultResult.Put(v)
}
