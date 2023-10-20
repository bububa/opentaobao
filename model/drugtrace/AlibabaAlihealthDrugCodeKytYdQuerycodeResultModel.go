package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel 结构体
type AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel struct {
	// 内层大对象
	Models []CodeFullInfoDto `json:"models,omitempty" xml:"models>code_full_info_dto,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 查询成功失败标记
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel() 从对象池中获取AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel
func GetAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel() *AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel {
	return poolAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel.Get().(*AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel 释放AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel
func ReleaseAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel(v *AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel) {
	v.Models = v.Models[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugCodeKytYdQuerycodeResultModel.Put(v)
}
