package wdk

import (
	"sync"
)

// AlibabaTclsAxIntegrationAccountImportApiResult 结构体
type AlibabaTclsAxIntegrationAccountImportApiResult struct {
	// 扩展信息
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 导入结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaTclsAxIntegrationAccountImportApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAxIntegrationAccountImportApiResult)
	},
}

// GetAlibabaTclsAxIntegrationAccountImportApiResult() 从对象池中获取AlibabaTclsAxIntegrationAccountImportApiResult
func GetAlibabaTclsAxIntegrationAccountImportApiResult() *AlibabaTclsAxIntegrationAccountImportApiResult {
	return poolAlibabaTclsAxIntegrationAccountImportApiResult.Get().(*AlibabaTclsAxIntegrationAccountImportApiResult)
}

// ReleaseAlibabaTclsAxIntegrationAccountImportApiResult 释放AlibabaTclsAxIntegrationAccountImportApiResult
func ReleaseAlibabaTclsAxIntegrationAccountImportApiResult(v *AlibabaTclsAxIntegrationAccountImportApiResult) {
	v.Ext = ""
	v.ErrorCode = ""
	v.Model = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaTclsAxIntegrationAccountImportApiResult.Put(v)
}
