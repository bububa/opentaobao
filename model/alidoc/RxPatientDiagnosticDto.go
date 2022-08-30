package alidoc

// RxPatientDiagnosticDto 结构体
type RxPatientDiagnosticDto struct {
	// 诊断list
	DiagnoseList []DiagnoseDto `json:"diagnose_list,omitempty" xml:"diagnose_list>diagnose_dto,omitempty"`
	// 过往病史
	MedicalHistory string `json:"medical_history,omitempty" xml:"medical_history,omitempty"`
	// 肝功能
	LiverFunction string `json:"liver_function,omitempty" xml:"liver_function,omitempty"`
	// 肾功能
	RenalFunction string `json:"renal_function,omitempty" xml:"renal_function,omitempty"`
	// 妊娠哺乳
	Pregnancy string `json:"pregnancy,omitempty" xml:"pregnancy,omitempty"`
	// 过敏史
	AllergyHistory string `json:"allergy_history,omitempty" xml:"allergy_history,omitempty"`
}
