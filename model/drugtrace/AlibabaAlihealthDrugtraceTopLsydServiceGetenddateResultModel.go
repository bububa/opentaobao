package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel struct {
	// 服务截止时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel
func GetAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel() *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel 释放AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel(v *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel) {
	v.EndDate = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel.Put(v)
}
