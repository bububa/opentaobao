package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceworkerCancelleaveResult 结构体
type AlibabaSscSupplyplatformServiceworkerCancelleaveResult struct {
	// 错误码
	DisplayCode string `json:"display_code,omitempty" xml:"display_code,omitempty"`
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceworkerCancelleaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerCancelleaveResult)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerCancelleaveResult() 从对象池中获取AlibabaSscSupplyplatformServiceworkerCancelleaveResult
func GetAlibabaSscSupplyplatformServiceworkerCancelleaveResult() *AlibabaSscSupplyplatformServiceworkerCancelleaveResult {
	return poolAlibabaSscSupplyplatformServiceworkerCancelleaveResult.Get().(*AlibabaSscSupplyplatformServiceworkerCancelleaveResult)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveResult 释放AlibabaSscSupplyplatformServiceworkerCancelleaveResult
func ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveResult(v *AlibabaSscSupplyplatformServiceworkerCancelleaveResult) {
	v.DisplayCode = ""
	v.DisplayMsg = ""
	v.IsSuccess = ""
	v.ErrorMsg = ""
	poolAlibabaSscSupplyplatformServiceworkerCancelleaveResult.Put(v)
}
