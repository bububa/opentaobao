package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServicestoreSaveResult 结构体
type AlibabaSscSupplyplatformServicestoreSaveResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回对象
	Value *ServiceStoreCreateDto `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServicestoreSaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicestoreSaveResult)
	},
}

// GetAlibabaSscSupplyplatformServicestoreSaveResult() 从对象池中获取AlibabaSscSupplyplatformServicestoreSaveResult
func GetAlibabaSscSupplyplatformServicestoreSaveResult() *AlibabaSscSupplyplatformServicestoreSaveResult {
	return poolAlibabaSscSupplyplatformServicestoreSaveResult.Get().(*AlibabaSscSupplyplatformServicestoreSaveResult)
}

// ReleaseAlibabaSscSupplyplatformServicestoreSaveResult 释放AlibabaSscSupplyplatformServicestoreSaveResult
func ReleaseAlibabaSscSupplyplatformServicestoreSaveResult(v *AlibabaSscSupplyplatformServicestoreSaveResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Value = nil
	v.Success = false
	poolAlibabaSscSupplyplatformServicestoreSaveResult.Put(v)
}
