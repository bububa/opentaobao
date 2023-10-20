package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflowModel 结构体
type AlibabaAlihealthDrugCodeKytQuerycodeflowModel struct {
	// 流向信息，如没有下游企业数据查看权限，部分数据会显示为空，发送授权邀请后可以正常显示
	CodeQueryFlows []CodeQueryFlows `json:"code_query_flows,omitempty" xml:"code_query_flows>code_query_flows,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 药品信息
	DrugInfo *DrugInfo `json:"drug_info,omitempty" xml:"drug_info,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeflowModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytQuerycodeflowModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeflowModel() 从对象池中获取AlibabaAlihealthDrugCodeKytQuerycodeflowModel
func GetAlibabaAlihealthDrugCodeKytQuerycodeflowModel() *AlibabaAlihealthDrugCodeKytQuerycodeflowModel {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeflowModel.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeflowModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowModel 释放AlibabaAlihealthDrugCodeKytQuerycodeflowModel
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowModel(v *AlibabaAlihealthDrugCodeKytQuerycodeflowModel) {
	v.CodeQueryFlows = v.CodeQueryFlows[:0]
	v.Code = ""
	v.ParentCode = ""
	v.DrugInfo = nil
	poolAlibabaAlihealthDrugCodeKytQuerycodeflowModel.Put(v)
}
