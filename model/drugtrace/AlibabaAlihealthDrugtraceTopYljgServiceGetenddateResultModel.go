package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel struct {
	// 服务截止时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel
func GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel() *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel {
	return poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel.Get().(*AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel 释放AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel
func ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel(v *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel) {
	v.EndDate = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel.Put(v)
}
