package alihealthpw

// ApplyInfo 结构体
type ApplyInfo struct {
	// 就诊信息
	TreatInfo *TreatInfo `json:"treat_info,omitempty" xml:"treat_info,omitempty"`
	// 填写人信息
	WriterInfo *WriterInfo `json:"writer_info,omitempty" xml:"writer_info,omitempty"`
	// 患者信息
	PatientInfo *PatientInfo `json:"patient_info,omitempty" xml:"patient_info,omitempty"`
}
