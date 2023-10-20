package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytYyApplycodeModel 结构体
type AlibabaAlihealthDrugCodeKytYyApplycodeModel struct {
	// 子码
	ChildrenCodeList []string `json:"children_code_list,omitempty" xml:"children_code_list>string,omitempty"`
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytYyApplycodeModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYyApplycodeModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytYyApplycodeModel() 从对象池中获取AlibabaAlihealthDrugCodeKytYyApplycodeModel
func GetAlibabaAlihealthDrugCodeKytYyApplycodeModel() *AlibabaAlihealthDrugCodeKytYyApplycodeModel {
	return poolAlibabaAlihealthDrugCodeKytYyApplycodeModel.Get().(*AlibabaAlihealthDrugCodeKytYyApplycodeModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeModel 释放AlibabaAlihealthDrugCodeKytYyApplycodeModel
func ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeModel(v *AlibabaAlihealthDrugCodeKytYyApplycodeModel) {
	v.ChildrenCodeList = v.ChildrenCodeList[:0]
	v.ParentCode = ""
	poolAlibabaAlihealthDrugCodeKytYyApplycodeModel.Put(v)
}
