package tmallservice

import (
	"sync"
)

// AlibabaSscSupplyplatformServiceworkerAvailableworkerResult 结构体
type AlibabaSscSupplyplatformServiceworkerAvailableworkerResult struct {
	// 可用工人列表
	Value []WorkerDto `json:"value,omitempty" xml:"value>worker_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统内部错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscSupplyplatformServiceworkerAvailableworkerResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerAvailableworkerResult)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerAvailableworkerResult() 从对象池中获取AlibabaSscSupplyplatformServiceworkerAvailableworkerResult
func GetAlibabaSscSupplyplatformServiceworkerAvailableworkerResult() *AlibabaSscSupplyplatformServiceworkerAvailableworkerResult {
	return poolAlibabaSscSupplyplatformServiceworkerAvailableworkerResult.Get().(*AlibabaSscSupplyplatformServiceworkerAvailableworkerResult)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerAvailableworkerResult 释放AlibabaSscSupplyplatformServiceworkerAvailableworkerResult
func ReleaseAlibabaSscSupplyplatformServiceworkerAvailableworkerResult(v *AlibabaSscSupplyplatformServiceworkerAvailableworkerResult) {
	v.Value = v.Value[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.DisplayMsg = ""
	v.Success = false
	poolAlibabaSscSupplyplatformServiceworkerAvailableworkerResult.Put(v)
}
