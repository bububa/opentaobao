package alihealthpw

// SNodeDto 结构体
type SNodeDto struct {
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 医疗费用补助
	HospitalAllowance string `json:"hospital_allowance,omitempty" xml:"hospital_allowance,omitempty"`
	// 医院收件电话
	HospitalEmsPhone string `json:"hospital_ems_phone,omitempty" xml:"hospital_ems_phone,omitempty"`
	// 是否转诊
	IsRefer string `json:"is_refer,omitempty" xml:"is_refer,omitempty"`
	// 资料链接
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// 医院地址
	HospitalAddress string `json:"hospital_address,omitempty" xml:"hospital_address,omitempty"`
	// 唯一编码
	UserUniqueCode string `json:"user_unique_code,omitempty" xml:"user_unique_code,omitempty"`
	// 审核时间
	ApplyAuditTime string `json:"apply_audit_time,omitempty" xml:"apply_audit_time,omitempty"`
	// 交通住宿打款日期
	FoundationPayDate string `json:"foundation_pay_date,omitempty" xml:"foundation_pay_date,omitempty"`
	// 审核意见
	CheckRemark string `json:"check_remark,omitempty" xml:"check_remark,omitempty"`
	// 医院电话
	HospitalPhone string `json:"hospital_phone,omitempty" xml:"hospital_phone,omitempty"`
	// 医院收件电话
	FoundationEmsPhone string `json:"foundation_ems_phone,omitempty" xml:"foundation_ems_phone,omitempty"`
	// 跳转url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 基金会收件地址
	FoundationEmsAddress string `json:"foundation_ems_address,omitempty" xml:"foundation_ems_address,omitempty"`
	// 交通住宿金额
	FoundationAllowance string `json:"foundation_allowance,omitempty" xml:"foundation_allowance,omitempty"`
	// 短信通知手机
	SmsPhone string `json:"sms_phone,omitempty" xml:"sms_phone,omitempty"`
	// 基金会收件人
	HospitalAddressee string `json:"hospital_addressee,omitempty" xml:"hospital_addressee,omitempty"`
	// 审核状态
	ApplyAuditStatus string `json:"apply_audit_status,omitempty" xml:"apply_audit_status,omitempty"`
	// 就诊医院
	TreatHospital string `json:"treat_hospital,omitempty" xml:"treat_hospital,omitempty"`
	// 基金会收件人
	FoundationAddressee string `json:"foundation_addressee,omitempty" xml:"foundation_addressee,omitempty"`
	// 医疗费用打款日期
	HospitalPayDate string `json:"hospital_pay_date,omitempty" xml:"hospital_pay_date,omitempty"`
	// 申请项目
	ProjectThirdId string `json:"project_third_id,omitempty" xml:"project_third_id,omitempty"`
	// 医院收件地址
	HospitalEmsAddress string `json:"hospital_ems_address,omitempty" xml:"hospital_ems_address,omitempty"`
	// 打款机构
	PaymentInstitution string `json:"payment_institution,omitempty" xml:"payment_institution,omitempty"`
}
