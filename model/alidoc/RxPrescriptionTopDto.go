package alidoc

// RxPrescriptionTopDto 结构体
type RxPrescriptionTopDto struct {
	// 药品列表
	DrugList []RxDrugTopDto `json:"drug_list,omitempty" xml:"drug_list>rx_drug_top_dto,omitempty"`
	// 患者问诊信息
	PatientDiagnostic *RxPatientDiagnosticTopDto `json:"patient_diagnostic,omitempty" xml:"patient_diagnostic,omitempty"`
	// 患者信息
	Patient *RxPatientTopDto `json:"patient,omitempty" xml:"patient,omitempty"`
	// 处方图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 处方创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 处方id
	RxId string `json:"rx_id,omitempty" xml:"rx_id,omitempty"`
	// 医生信息
	Doctor *RxDoctorTopDto `json:"doctor,omitempty" xml:"doctor,omitempty"`
}
