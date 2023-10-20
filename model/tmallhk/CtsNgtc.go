package tmallhk

import (
	"sync"
)

// CtsNgtc 结构体
type CtsNgtc struct {
	// ngtc证书编号
	ReportNo string `json:"report_no,omitempty" xml:"report_no,omitempty"`
	// ngtc证书防伪码
	ReportVerifyNo string `json:"report_verify_no,omitempty" xml:"report_verify_no,omitempty"`
}

var poolCtsNgtc = sync.Pool{
	New: func() any {
		return new(CtsNgtc)
	},
}

// GetCtsNgtc() 从对象池中获取CtsNgtc
func GetCtsNgtc() *CtsNgtc {
	return poolCtsNgtc.Get().(*CtsNgtc)
}

// ReleaseCtsNgtc 释放CtsNgtc
func ReleaseCtsNgtc(v *CtsNgtc) {
	v.ReportNo = ""
	v.ReportVerifyNo = ""
	poolCtsNgtc.Put(v)
}
