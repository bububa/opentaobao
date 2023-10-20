package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult 结构体
type AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult struct {
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult() 从对象池中获取AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult
func GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult() *AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult {
	return poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult.Get().(*AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult 释放AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult
func ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult(v *AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreResult.Put(v)
}
