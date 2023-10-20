package guoguo

import (
	"sync"
)

// CainiaoGuoguoBackupGraborderSubmitmailnoResult 结构体
type CainiaoGuoguoBackupGraborderSubmitmailnoResult struct {
	// 1
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 返回的状态描述
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 数据对象
	Data *BackupOrderDo `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoGuoguoBackupGraborderSubmitmailnoResult = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoBackupGraborderSubmitmailnoResult)
	},
}

// GetCainiaoGuoguoBackupGraborderSubmitmailnoResult() 从对象池中获取CainiaoGuoguoBackupGraborderSubmitmailnoResult
func GetCainiaoGuoguoBackupGraborderSubmitmailnoResult() *CainiaoGuoguoBackupGraborderSubmitmailnoResult {
	return poolCainiaoGuoguoBackupGraborderSubmitmailnoResult.Get().(*CainiaoGuoguoBackupGraborderSubmitmailnoResult)
}

// ReleaseCainiaoGuoguoBackupGraborderSubmitmailnoResult 释放CainiaoGuoguoBackupGraborderSubmitmailnoResult
func ReleaseCainiaoGuoguoBackupGraborderSubmitmailnoResult(v *CainiaoGuoguoBackupGraborderSubmitmailnoResult) {
	v.StatusCode = ""
	v.StatusMessage = ""
	v.Data = nil
	v.Success = false
	poolCainiaoGuoguoBackupGraborderSubmitmailnoResult.Put(v)
}
