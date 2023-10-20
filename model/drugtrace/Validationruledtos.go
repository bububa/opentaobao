package drugtrace

import (
	"sync"
)

// Validationruledtos 结构体
type Validationruledtos struct {
	// 校验的字段名称，如expire_date（有效期）、sdc_code（本位码）、prepn_spec（制剂规格）
	FieldName string `json:"field_name,omitempty" xml:"field_name,omitempty"`
	// 校验结果Code,  0:通过；1:不通过
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 校验结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
}

var poolValidationruledtos = sync.Pool{
	New: func() any {
		return new(Validationruledtos)
	},
}

// GetValidationruledtos() 从对象池中获取Validationruledtos
func GetValidationruledtos() *Validationruledtos {
	return poolValidationruledtos.Get().(*Validationruledtos)
}

// ReleaseValidationruledtos 释放Validationruledtos
func ReleaseValidationruledtos(v *Validationruledtos) {
	v.FieldName = ""
	v.ResultCode = ""
	v.ResultDesc = ""
	poolValidationruledtos.Put(v)
}
