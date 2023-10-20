package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyPutpackageResultModel 结构体
type AlibabaAlihealthDrugKytScqyPutpackageResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 内层大对象
	Model *NewCodeMovePackagingResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 查询成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyPutpackageResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyPutpackageResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyPutpackageResultModel
func GetAlibabaAlihealthDrugKytScqyPutpackageResultModel() *AlibabaAlihealthDrugKytScqyPutpackageResultModel {
	return poolAlibabaAlihealthDrugKytScqyPutpackageResultModel.Get().(*AlibabaAlihealthDrugKytScqyPutpackageResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageResultModel 释放AlibabaAlihealthDrugKytScqyPutpackageResultModel
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageResultModel(v *AlibabaAlihealthDrugKytScqyPutpackageResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytScqyPutpackageResultModel.Put(v)
}
