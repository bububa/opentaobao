package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel 结构体
type AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel struct {
	// 内层大对象
	Models []CodeFullInfoDto `json:"models,omitempty" xml:"models>code_full_info_dto,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 查询成功失败标记
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel() 从对象池中获取AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel
func GetAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel() *AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel {
	return poolAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel.Get().(*AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel 释放AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel
func ReleaseAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel(v *AlibabaAlihealthDrugCodeKytDrQuerycodeResultModel) {
	v.Models = v.Models[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugCodeKytDrQuerycodeResultModel.Put(v)
}
