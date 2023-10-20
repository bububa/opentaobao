package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel 结构体
type AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 内层大对象
	Model *CodeMovePackagingSecondDto `json:"model,omitempty" xml:"model,omitempty"`
	// 查询成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel)
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel() 从对象池中获取AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel
func GetAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel() *AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel {
	return poolAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel.Get().(*AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel 释放AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel(v *AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytScqyPutpackageunbindResultModel.Put(v)
}
