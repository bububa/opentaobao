package tmallhk

import (
	"sync"
)

// CertifyQueryResult 结构体
type CertifyQueryResult struct {
	// 错误原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 清关对象
	Module *ConsigneeCertifyInfo `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCertifyQueryResult = sync.Pool{
	New: func() any {
		return new(CertifyQueryResult)
	},
}

// GetCertifyQueryResult() 从对象池中获取CertifyQueryResult
func GetCertifyQueryResult() *CertifyQueryResult {
	return poolCertifyQueryResult.Get().(*CertifyQueryResult)
}

// ReleaseCertifyQueryResult 释放CertifyQueryResult
func ReleaseCertifyQueryResult(v *CertifyQueryResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Module = nil
	v.Success = false
	poolCertifyQueryResult.Put(v)
}
