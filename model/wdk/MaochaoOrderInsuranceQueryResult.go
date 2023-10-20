package wdk

import (
	"sync"
)

// MaochaoOrderInsuranceQueryResult 结构体
type MaochaoOrderInsuranceQueryResult struct {
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回结果
	Model *InsuranceOrder `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMaochaoOrderInsuranceQueryResult = sync.Pool{
	New: func() any {
		return new(MaochaoOrderInsuranceQueryResult)
	},
}

// GetMaochaoOrderInsuranceQueryResult() 从对象池中获取MaochaoOrderInsuranceQueryResult
func GetMaochaoOrderInsuranceQueryResult() *MaochaoOrderInsuranceQueryResult {
	return poolMaochaoOrderInsuranceQueryResult.Get().(*MaochaoOrderInsuranceQueryResult)
}

// ReleaseMaochaoOrderInsuranceQueryResult 释放MaochaoOrderInsuranceQueryResult
func ReleaseMaochaoOrderInsuranceQueryResult(v *MaochaoOrderInsuranceQueryResult) {
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.Model = nil
	v.Success = false
	poolMaochaoOrderInsuranceQueryResult.Put(v)
}
