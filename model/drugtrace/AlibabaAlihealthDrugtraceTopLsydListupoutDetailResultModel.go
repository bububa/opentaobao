package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel struct {
	// 提示信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 提示信息内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 最外层对象
	Model *BillUpOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel
func GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel() *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel 释放AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel(v *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel.Put(v)
}
