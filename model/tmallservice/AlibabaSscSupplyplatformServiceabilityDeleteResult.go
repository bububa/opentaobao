package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceabilityDeleteResult 结构体
type AlibabaSscSupplyplatformServiceabilityDeleteResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceabilityDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceabilityDeleteResult)
	},
}

// GetAlibabaSscSupplyplatformServiceabilityDeleteResult() 从对象池中获取AlibabaSscSupplyplatformServiceabilityDeleteResult
func GetAlibabaSscSupplyplatformServiceabilityDeleteResult() *AlibabaSscSupplyplatformServiceabilityDeleteResult {
	return poolAlibabaSscSupplyplatformServiceabilityDeleteResult.Get().(*AlibabaSscSupplyplatformServiceabilityDeleteResult)
}

// ReleaseAlibabaSscSupplyplatformServiceabilityDeleteResult 释放AlibabaSscSupplyplatformServiceabilityDeleteResult
func ReleaseAlibabaSscSupplyplatformServiceabilityDeleteResult(v *AlibabaSscSupplyplatformServiceabilityDeleteResult) {
	v.DisplayMsg = ""
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServiceabilityDeleteResult.Put(v)
}
