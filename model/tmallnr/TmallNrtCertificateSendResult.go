package tmallnr

import (
	"sync"
)

// TmallNrtCertificateSendResult 结构体
type TmallNrtCertificateSendResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// model
	Model *NrtCertificateInstanceResponseDto `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTmallNrtCertificateSendResult = sync.Pool{
	New: func() any {
		return new(TmallNrtCertificateSendResult)
	},
}

// GetTmallNrtCertificateSendResult() 从对象池中获取TmallNrtCertificateSendResult
func GetTmallNrtCertificateSendResult() *TmallNrtCertificateSendResult {
	return poolTmallNrtCertificateSendResult.Get().(*TmallNrtCertificateSendResult)
}

// ReleaseTmallNrtCertificateSendResult 释放TmallNrtCertificateSendResult
func ReleaseTmallNrtCertificateSendResult(v *TmallNrtCertificateSendResult) {
	v.Message = ""
	v.Code = ""
	v.Model = nil
	poolTmallNrtCertificateSendResult.Put(v)
}
