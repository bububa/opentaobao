package alidoc

// PrescriptionSearchResultDto 结构体
type PrescriptionSearchResultDto struct {
	// 药品列表
	DrugList []DrugDto `json:"drug_list,omitempty" xml:"drug_list>drug_dto,omitempty"`
	// 处方id
	RxId string `json:"rx_id,omitempty" xml:"rx_id,omitempty"`
	// 处方创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 处方图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 医生信息
	Doctor *RxDoctorDto `json:"doctor,omitempty" xml:"doctor,omitempty"`
	// 患者信息
	Patient *RxPatientDto `json:"patient,omitempty" xml:"patient,omitempty"`
	// 患者问诊信息
	PatientDiagnostic *RxPatientDiagnosticDto `json:"patient_diagnostic,omitempty" xml:"patient_diagnostic,omitempty"`
}
