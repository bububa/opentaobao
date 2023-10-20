package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceworkerSaveResult 结构体
type AlibabaSscSupplyplatformServiceworkerSaveResult struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 展示信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceworkerSaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerSaveResult)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerSaveResult() 从对象池中获取AlibabaSscSupplyplatformServiceworkerSaveResult
func GetAlibabaSscSupplyplatformServiceworkerSaveResult() *AlibabaSscSupplyplatformServiceworkerSaveResult {
	return poolAlibabaSscSupplyplatformServiceworkerSaveResult.Get().(*AlibabaSscSupplyplatformServiceworkerSaveResult)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerSaveResult 释放AlibabaSscSupplyplatformServiceworkerSaveResult
func ReleaseAlibabaSscSupplyplatformServiceworkerSaveResult(v *AlibabaSscSupplyplatformServiceworkerSaveResult) {
	v.ErrorCode = ""
	v.DisplayMsg = ""
	v.Success = ""
	v.ErrorMsg = ""
	poolAlibabaSscSupplyplatformServiceworkerSaveResult.Put(v)
}
