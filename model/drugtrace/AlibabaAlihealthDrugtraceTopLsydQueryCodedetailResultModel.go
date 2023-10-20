package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel struct {
	// 内层大对象
	Models []CodeFullInfoDto `json:"models,omitempty" xml:"models>code_full_info_dto,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 查询成功失败标记
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel
func GetAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel() *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel 释放AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel(v *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel) {
	v.Models = v.Models[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel.Put(v)
}
