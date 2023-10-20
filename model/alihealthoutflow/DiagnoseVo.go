package alihealthoutflow

import (
	"sync"
)

// DiagnoseVo 结构体
type DiagnoseVo struct {
	// 诊断编码
	DiagnoseCode string `json:"diagnose_code,omitempty" xml:"diagnose_code,omitempty"`
	// 诊断名
	DiagnoseName string `json:"diagnose_name,omitempty" xml:"diagnose_name,omitempty"`
}

var poolDiagnoseVo = sync.Pool{
	New: func() any {
		return new(DiagnoseVo)
	},
}

// GetDiagnoseVo() 从对象池中获取DiagnoseVo
func GetDiagnoseVo() *DiagnoseVo {
	return poolDiagnoseVo.Get().(*DiagnoseVo)
}

// ReleaseDiagnoseVo 释放DiagnoseVo
func ReleaseDiagnoseVo(v *DiagnoseVo) {
	v.DiagnoseCode = ""
	v.DiagnoseName = ""
	poolDiagnoseVo.Put(v)
}
