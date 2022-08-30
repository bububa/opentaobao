package tmallnr

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
