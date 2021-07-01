package alihealthoutflow

// Diagnose 结构体
type Diagnose struct {
	// 诊断编码-his(his、icd10二选一)
	DiagnoseCode string `json:"diagnose_code,omitempty" xml:"diagnose_code,omitempty"`
	// 诊断编码-his(his、icd10二选一)
	DiagnoseName string `json:"diagnose_name,omitempty" xml:"diagnose_name,omitempty"`
	// icd9/icd10名称(his、icd10二选一) - 纳里必传
	IcdName string `json:"icd_name,omitempty" xml:"icd_name,omitempty"`
	// icd9/icd10编码(his、icd10二选一) - 纳里必传
	IcdCode string `json:"icd_code,omitempty" xml:"icd_code,omitempty"`
}
