package tmallnr

import (
	"sync"
)

// NrtUpdateActivityReq 结构体
type NrtUpdateActivityReq struct {
	// 权利说明
	CertificateRights string `json:"certificate_rights,omitempty" xml:"certificate_rights,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 页面ID
	PageId int64 `json:"page_id,omitempty" xml:"page_id,omitempty"`
	// 员工ID
	EmployeeId int64 `json:"employee_id,omitempty" xml:"employee_id,omitempty"`
	// 是否需要电子凭证
	NeedCertificate bool `json:"need_certificate,omitempty" xml:"need_certificate,omitempty"`
}

var poolNrtUpdateActivityReq = sync.Pool{
	New: func() any {
		return new(NrtUpdateActivityReq)
	},
}

// GetNrtUpdateActivityReq() 从对象池中获取NrtUpdateActivityReq
func GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
	return poolNrtUpdateActivityReq.Get().(*NrtUpdateActivityReq)
}

// ReleaseNrtUpdateActivityReq 释放NrtUpdateActivityReq
func ReleaseNrtUpdateActivityReq(v *NrtUpdateActivityReq) {
	v.CertificateRights = ""
	v.ActivityId = 0
	v.PageId = 0
	v.EmployeeId = 0
	v.NeedCertificate = false
	poolNrtUpdateActivityReq.Put(v)
}
