package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceabilitySaveResult 结构体
type AlibabaSscSupplyplatformServiceabilitySaveResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceabilitySaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceabilitySaveResult)
	},
}

// GetAlibabaSscSupplyplatformServiceabilitySaveResult() 从对象池中获取AlibabaSscSupplyplatformServiceabilitySaveResult
func GetAlibabaSscSupplyplatformServiceabilitySaveResult() *AlibabaSscSupplyplatformServiceabilitySaveResult {
	return poolAlibabaSscSupplyplatformServiceabilitySaveResult.Get().(*AlibabaSscSupplyplatformServiceabilitySaveResult)
}

// ReleaseAlibabaSscSupplyplatformServiceabilitySaveResult 释放AlibabaSscSupplyplatformServiceabilitySaveResult
func ReleaseAlibabaSscSupplyplatformServiceabilitySaveResult(v *AlibabaSscSupplyplatformServiceabilitySaveResult) {
	v.DisplayMsg = ""
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServiceabilitySaveResult.Put(v)
}
