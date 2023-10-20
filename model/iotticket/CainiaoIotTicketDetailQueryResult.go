package iotticket

import (
	"sync"
)

// CainiaoIotTicketDetailQueryResult 结构体
type CainiaoIotTicketDetailQueryResult struct {
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果数据
	Data *CainiaoIotTicketDetailQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketDetailQueryResult = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketDetailQueryResult)
	},
}

// GetCainiaoIotTicketDetailQueryResult() 从对象池中获取CainiaoIotTicketDetailQueryResult
func GetCainiaoIotTicketDetailQueryResult() *CainiaoIotTicketDetailQueryResult {
	return poolCainiaoIotTicketDetailQueryResult.Get().(*CainiaoIotTicketDetailQueryResult)
}

// ReleaseCainiaoIotTicketDetailQueryResult 释放CainiaoIotTicketDetailQueryResult
func ReleaseCainiaoIotTicketDetailQueryResult(v *CainiaoIotTicketDetailQueryResult) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolCainiaoIotTicketDetailQueryResult.Put(v)
}
