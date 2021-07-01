package alihealthoutflow

// DiagnoseVo 结构体
type DiagnoseVo struct {
	// 诊断编码
	DiagnoseCode string `json:"diagnose_code,omitempty" xml:"diagnose_code,omitempty"`
	// 诊断名
	DiagnoseName string `json:"diagnose_name,omitempty" xml:"diagnose_name,omitempty"`
}
