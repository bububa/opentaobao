package alihealthoutflow

// PrescriptionMoreDataResponse 结构体
type PrescriptionMoreDataResponse struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 过期时间
	ExpiredTime string `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	// 费用类别
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 药品列表
	DrugList []DrugDto `json:"drug_list,omitempty" xml:"drug_list>drug_dto,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 处方类型
	RxType string `json:"rx_type,omitempty" xml:"rx_type,omitempty"`
	// 处方发送时间
	SendPrescriptionTime string `json:"send_prescription_time,omitempty" xml:"send_prescription_time,omitempty"`
	// 平台code
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 处方号
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 药品类别
	DrugType string `json:"drug_type,omitempty" xml:"drug_type,omitempty"`
	// 服务类别
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 诊断
	DiagnoseName string `json:"diagnose_name,omitempty" xml:"diagnose_name,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 年龄
	Age string `json:"age,omitempty" xml:"age,omitempty"`
	// 医生姓名
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 医院名字
	HospitalName string `json:"hospital_name,omitempty" xml:"hospital_name,omitempty"`
	// 药品数量
	RxCount string `json:"rx_count,omitempty" xml:"rx_count,omitempty"`
	// 外配处方备案号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 患者电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 图片地址
	RxImageUrl string `json:"rx_image_url,omitempty" xml:"rx_image_url,omitempty"`
}
