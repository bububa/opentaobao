package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeActivityCustomerSaveResult 结构体
type AlibabaAlihouseNewhomeActivityCustomerSaveResult struct {
	// 处理失败客户集合
	Data []ActivityCustomerErrHandlerResultDto `json:"data,omitempty" xml:"data>activity_customer_err_handler_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeActivityCustomerSaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivityCustomerSaveResult)
	},
}

// GetAlibabaAlihouseNewhomeActivityCustomerSaveResult() 从对象池中获取AlibabaAlihouseNewhomeActivityCustomerSaveResult
func GetAlibabaAlihouseNewhomeActivityCustomerSaveResult() *AlibabaAlihouseNewhomeActivityCustomerSaveResult {
	return poolAlibabaAlihouseNewhomeActivityCustomerSaveResult.Get().(*AlibabaAlihouseNewhomeActivityCustomerSaveResult)
}

// ReleaseAlibabaAlihouseNewhomeActivityCustomerSaveResult 释放AlibabaAlihouseNewhomeActivityCustomerSaveResult
func ReleaseAlibabaAlihouseNewhomeActivityCustomerSaveResult(v *AlibabaAlihouseNewhomeActivityCustomerSaveResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeActivityCustomerSaveResult.Put(v)
}
