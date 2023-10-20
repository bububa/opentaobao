package eticket

import (
	"sync"
)

// QueryTbMaCallbackResp 结构体
type QueryTbMaCallbackResp struct {
	// certificateDTO
	Certificate *CertificateDto `json:"certificate,omitempty" xml:"certificate,omitempty"`
}

var poolQueryTbMaCallbackResp = sync.Pool{
	New: func() any {
		return new(QueryTbMaCallbackResp)
	},
}

// GetQueryTbMaCallbackResp() 从对象池中获取QueryTbMaCallbackResp
func GetQueryTbMaCallbackResp() *QueryTbMaCallbackResp {
	return poolQueryTbMaCallbackResp.Get().(*QueryTbMaCallbackResp)
}

// ReleaseQueryTbMaCallbackResp 释放QueryTbMaCallbackResp
func ReleaseQueryTbMaCallbackResp(v *QueryTbMaCallbackResp) {
	v.Certificate = nil
	poolQueryTbMaCallbackResp.Put(v)
}
