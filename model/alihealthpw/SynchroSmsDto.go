package alihealthpw

// SynchroSmsDto 结构体
type SynchroSmsDto struct {
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 短信通知手机号
	SmsPhone string `json:"sms_phone,omitempty" xml:"sms_phone,omitempty"`
	// 就诊医院
	TreatHospital string `json:"treat_hospital,omitempty" xml:"treat_hospital,omitempty"`
	// 唯一编码
	UserUniqueCode string `json:"user_unique_code,omitempty" xml:"user_unique_code,omitempty"`
	// 申请项目
	ProjectThirdId string `json:"project_third_id,omitempty" xml:"project_third_id,omitempty"`
	// 医院收件人
	HospitalAddressee string `json:"hospital_addressee,omitempty" xml:"hospital_addressee,omitempty"`
	// 医院收件电话
	HospitalEmsPhone string `json:"hospital_ems_phone,omitempty" xml:"hospital_ems_phone,omitempty"`
	// 医院收件地址
	HospitalEmsAddress string `json:"hospital_ems_address,omitempty" xml:"hospital_ems_address,omitempty"`
	// 医院地址
	HospitalAddress string `json:"hospital_address,omitempty" xml:"hospital_address,omitempty"`
	// 医院电话
	HospitalPhone string `json:"hospital_phone,omitempty" xml:"hospital_phone,omitempty"`
}
