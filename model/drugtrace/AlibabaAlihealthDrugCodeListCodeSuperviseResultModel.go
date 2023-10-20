package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeListCodeSuperviseResultModel 结构体
type AlibabaAlihealthDrugCodeListCodeSuperviseResultModel struct {
	// 内层大对象
	Models []CodeFullInfoDto `json:"models,omitempty" xml:"models>code_full_info_dto,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 查询成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeListCodeSuperviseResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeListCodeSuperviseResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeListCodeSuperviseResultModel() 从对象池中获取AlibabaAlihealthDrugCodeListCodeSuperviseResultModel
func GetAlibabaAlihealthDrugCodeListCodeSuperviseResultModel() *AlibabaAlihealthDrugCodeListCodeSuperviseResultModel {
	return poolAlibabaAlihealthDrugCodeListCodeSuperviseResultModel.Get().(*AlibabaAlihealthDrugCodeListCodeSuperviseResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeListCodeSuperviseResultModel 释放AlibabaAlihealthDrugCodeListCodeSuperviseResultModel
func ReleaseAlibabaAlihealthDrugCodeListCodeSuperviseResultModel(v *AlibabaAlihealthDrugCodeListCodeSuperviseResultModel) {
	v.Models = v.Models[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthDrugCodeListCodeSuperviseResultModel.Put(v)
}
