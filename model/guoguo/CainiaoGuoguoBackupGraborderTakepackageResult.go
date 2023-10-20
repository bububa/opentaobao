package guoguo

import (
	"sync"
)

// CainiaoGuoguoBackupGraborderTakepackageResult 结构体
type CainiaoGuoguoBackupGraborderTakepackageResult struct {
	// 调用状态
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 调用状态码
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 封装返回的数据
	Data *BackupOrderDo `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoGuoguoBackupGraborderTakepackageResult = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoBackupGraborderTakepackageResult)
	},
}

// GetCainiaoGuoguoBackupGraborderTakepackageResult() 从对象池中获取CainiaoGuoguoBackupGraborderTakepackageResult
func GetCainiaoGuoguoBackupGraborderTakepackageResult() *CainiaoGuoguoBackupGraborderTakepackageResult {
	return poolCainiaoGuoguoBackupGraborderTakepackageResult.Get().(*CainiaoGuoguoBackupGraborderTakepackageResult)
}

// ReleaseCainiaoGuoguoBackupGraborderTakepackageResult 释放CainiaoGuoguoBackupGraborderTakepackageResult
func ReleaseCainiaoGuoguoBackupGraborderTakepackageResult(v *CainiaoGuoguoBackupGraborderTakepackageResult) {
	v.StatusCode = ""
	v.StatusMessage = ""
	v.Data = nil
	v.Success = false
	poolCainiaoGuoguoBackupGraborderTakepackageResult.Put(v)
}
