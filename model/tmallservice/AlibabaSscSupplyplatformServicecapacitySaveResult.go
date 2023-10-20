package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServicecapacitySaveResult 结构体
type AlibabaSscSupplyplatformServicecapacitySaveResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServicecapacitySaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicecapacitySaveResult)
	},
}

// GetAlibabaSscSupplyplatformServicecapacitySaveResult() 从对象池中获取AlibabaSscSupplyplatformServicecapacitySaveResult
func GetAlibabaSscSupplyplatformServicecapacitySaveResult() *AlibabaSscSupplyplatformServicecapacitySaveResult {
	return poolAlibabaSscSupplyplatformServicecapacitySaveResult.Get().(*AlibabaSscSupplyplatformServicecapacitySaveResult)
}

// ReleaseAlibabaSscSupplyplatformServicecapacitySaveResult 释放AlibabaSscSupplyplatformServicecapacitySaveResult
func ReleaseAlibabaSscSupplyplatformServicecapacitySaveResult(v *AlibabaSscSupplyplatformServicecapacitySaveResult) {
	v.DisplayMsg = ""
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServicecapacitySaveResult.Put(v)
}
