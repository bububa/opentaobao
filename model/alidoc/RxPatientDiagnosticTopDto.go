package alidoc

// RxPatientDiagnosticTopDto 结构体
type RxPatientDiagnosticTopDto struct {
	// 诊断list
	DiagnoseList []RxDiagnoseTopDto `json:"diagnose_list,omitempty" xml:"diagnose_list>rx_diagnose_top_dto,omitempty"`
	// 过敏史
	AllergyHistory string `json:"allergy_history,omitempty" xml:"allergy_history,omitempty"`
	// 妊娠哺乳
	Pregnancy string `json:"pregnancy,omitempty" xml:"pregnancy,omitempty"`
	// 肾功能
	RenalFunction string `json:"renal_function,omitempty" xml:"renal_function,omitempty"`
	// 肝功能
	LiverFunction string `json:"liver_function,omitempty" xml:"liver_function,omitempty"`
	// 过往病史
	MedicalHistory string `json:"medical_history,omitempty" xml:"medical_history,omitempty"`
}
