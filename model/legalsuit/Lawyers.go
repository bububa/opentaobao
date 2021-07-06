package legalsuit

// Lawyers 结构体
type Lawyers struct {
	// 附件列表
	AttachmentList []Attachmentlist `json:"attachment_list,omitempty" xml:"attachment_list>attachmentlist,omitempty"`
	// 律师策略
	ResponseStrategy string `json:"response_strategy,omitempty" xml:"response_strategy,omitempty"`
	// 律师费用
	ChargeThisCase string `json:"charge_this_case,omitempty" xml:"charge_this_case,omitempty"`
	// 执业证编号
	ProfessionalCertNum string `json:"professional_cert_num,omitempty" xml:"professional_cert_num,omitempty"`
	// 供应商
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 律所编号
	BusinessLicenseNum string `json:"business_license_num,omitempty" xml:"business_license_num,omitempty"`
	// 律师名称
	LawyerName string `json:"lawyer_name,omitempty" xml:"lawyer_name,omitempty"`
	// 律所名称
	LawyerFirmName string `json:"lawyer_firm_name,omitempty" xml:"lawyer_firm_name,omitempty"`
	// 附件数量
	AttachmentCount int64 `json:"attachment_count,omitempty" xml:"attachment_count,omitempty"`
	// 律师ID
	LawyerId int64 `json:"lawyer_id,omitempty" xml:"lawyer_id,omitempty"`
}
