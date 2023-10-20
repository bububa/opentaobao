package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceworkerRegisterResult 结构体
type AlibabaSscSupplyplatformServiceworkerRegisterResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 结果model
	Value *WorkerCreateDto `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceworkerRegisterResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerRegisterResult)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerRegisterResult() 从对象池中获取AlibabaSscSupplyplatformServiceworkerRegisterResult
func GetAlibabaSscSupplyplatformServiceworkerRegisterResult() *AlibabaSscSupplyplatformServiceworkerRegisterResult {
	return poolAlibabaSscSupplyplatformServiceworkerRegisterResult.Get().(*AlibabaSscSupplyplatformServiceworkerRegisterResult)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerRegisterResult 释放AlibabaSscSupplyplatformServiceworkerRegisterResult
func ReleaseAlibabaSscSupplyplatformServiceworkerRegisterResult(v *AlibabaSscSupplyplatformServiceworkerRegisterResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.DisplayMsg = ""
	v.Value = nil
	v.Success = false
	poolAlibabaSscSupplyplatformServiceworkerRegisterResult.Put(v)
}
