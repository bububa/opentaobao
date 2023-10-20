package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServicecapacityDeleteResult 结构体
type AlibabaSscSupplyplatformServicecapacityDeleteResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServicecapacityDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicecapacityDeleteResult)
	},
}

// GetAlibabaSscSupplyplatformServicecapacityDeleteResult() 从对象池中获取AlibabaSscSupplyplatformServicecapacityDeleteResult
func GetAlibabaSscSupplyplatformServicecapacityDeleteResult() *AlibabaSscSupplyplatformServicecapacityDeleteResult {
	return poolAlibabaSscSupplyplatformServicecapacityDeleteResult.Get().(*AlibabaSscSupplyplatformServicecapacityDeleteResult)
}

// ReleaseAlibabaSscSupplyplatformServicecapacityDeleteResult 释放AlibabaSscSupplyplatformServicecapacityDeleteResult
func ReleaseAlibabaSscSupplyplatformServicecapacityDeleteResult(v *AlibabaSscSupplyplatformServicecapacityDeleteResult) {
	v.DisplayMsg = ""
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServicecapacityDeleteResult.Put(v)
}
