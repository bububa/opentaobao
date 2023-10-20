package logistic

import (
	"sync"
)

// TmsCourierInfoRequest 结构体
type TmsCourierInfoRequest struct {
	// 小件员名称
	CourierName string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
	// 小件员电话号码
	CourierMobile string `json:"courier_mobile,omitempty" xml:"courier_mobile,omitempty"`
	// 小件员编码
	CourierNo string `json:"courier_no,omitempty" xml:"courier_no,omitempty"`
	// 小件员所属的网点名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 小件员所属的网点编码
	SiteCode string `json:"site_code,omitempty" xml:"site_code,omitempty"`
}

var poolTmsCourierInfoRequest = sync.Pool{
	New: func() any {
		return new(TmsCourierInfoRequest)
	},
}

// GetTmsCourierInfoRequest() 从对象池中获取TmsCourierInfoRequest
func GetTmsCourierInfoRequest() *TmsCourierInfoRequest {
	return poolTmsCourierInfoRequest.Get().(*TmsCourierInfoRequest)
}

// ReleaseTmsCourierInfoRequest 释放TmsCourierInfoRequest
func ReleaseTmsCourierInfoRequest(v *TmsCourierInfoRequest) {
	v.CourierName = ""
	v.CourierMobile = ""
	v.CourierNo = ""
	v.SiteName = ""
	v.SiteCode = ""
	poolTmsCourierInfoRequest.Put(v)
}
