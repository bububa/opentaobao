package bus

import (
	"sync"
)

// InsurancePropertyVo 结构体
type InsurancePropertyVo struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 数据
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolInsurancePropertyVo = sync.Pool{
	New: func() any {
		return new(InsurancePropertyVo)
	},
}

// GetInsurancePropertyVo() 从对象池中获取InsurancePropertyVo
func GetInsurancePropertyVo() *InsurancePropertyVo {
	return poolInsurancePropertyVo.Get().(*InsurancePropertyVo)
}

// ReleaseInsurancePropertyVo 释放InsurancePropertyVo
func ReleaseInsurancePropertyVo(v *InsurancePropertyVo) {
	v.Name = ""
	v.Type = ""
	v.Value = ""
	v.Key = ""
	poolInsurancePropertyVo.Put(v)
}
