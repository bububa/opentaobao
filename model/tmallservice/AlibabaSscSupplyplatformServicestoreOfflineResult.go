package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServicestoreOfflineResult 结构体
type AlibabaSscSupplyplatformServicestoreOfflineResult struct {
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServicestoreOfflineResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicestoreOfflineResult)
	},
}

// GetAlibabaSscSupplyplatformServicestoreOfflineResult() 从对象池中获取AlibabaSscSupplyplatformServicestoreOfflineResult
func GetAlibabaSscSupplyplatformServicestoreOfflineResult() *AlibabaSscSupplyplatformServicestoreOfflineResult {
	return poolAlibabaSscSupplyplatformServicestoreOfflineResult.Get().(*AlibabaSscSupplyplatformServicestoreOfflineResult)
}

// ReleaseAlibabaSscSupplyplatformServicestoreOfflineResult 释放AlibabaSscSupplyplatformServicestoreOfflineResult
func ReleaseAlibabaSscSupplyplatformServicestoreOfflineResult(v *AlibabaSscSupplyplatformServicestoreOfflineResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServicestoreOfflineResult.Put(v)
}
