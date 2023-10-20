package tmallnr

import (
	"sync"
)

// TmallNrtCertificateQueryResult 结构体
type TmallNrtCertificateQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// model
	Model *PageData `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTmallNrtCertificateQueryResult = sync.Pool{
	New: func() any {
		return new(TmallNrtCertificateQueryResult)
	},
}

// GetTmallNrtCertificateQueryResult() 从对象池中获取TmallNrtCertificateQueryResult
func GetTmallNrtCertificateQueryResult() *TmallNrtCertificateQueryResult {
	return poolTmallNrtCertificateQueryResult.Get().(*TmallNrtCertificateQueryResult)
}

// ReleaseTmallNrtCertificateQueryResult 释放TmallNrtCertificateQueryResult
func ReleaseTmallNrtCertificateQueryResult(v *TmallNrtCertificateQueryResult) {
	v.Message = ""
	v.Code = ""
	v.Model = nil
	poolTmallNrtCertificateQueryResult.Put(v)
}
