package drugtrace

import (
	"sync"
)

// DrugCode 结构体
type DrugCode struct {
	// 包装规格
	PkgSpecList []string `json:"pkg_spec_list,omitempty" xml:"pkg_spec_list>string,omitempty"`
	// 是否有码
	HasCode bool `json:"has_code,omitempty" xml:"has_code,omitempty"`
}

var poolDrugCode = sync.Pool{
	New: func() any {
		return new(DrugCode)
	},
}

// GetDrugCode() 从对象池中获取DrugCode
func GetDrugCode() *DrugCode {
	return poolDrugCode.Get().(*DrugCode)
}

// ReleaseDrugCode 释放DrugCode
func ReleaseDrugCode(v *DrugCode) {
	v.PkgSpecList = v.PkgSpecList[:0]
	v.HasCode = false
	poolDrugCode.Put(v)
}
