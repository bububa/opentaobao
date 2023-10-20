package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQuerydruginfoResultModel 结构体
type AlibabaAlihealthDrugKytQuerydruginfoResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回业务对象
	Model *CodeQueryDrugInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQuerydruginfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerydruginfoResultModel)
	},
}

// GetAlibabaAlihealthDrugKytQuerydruginfoResultModel() 从对象池中获取AlibabaAlihealthDrugKytQuerydruginfoResultModel
func GetAlibabaAlihealthDrugKytQuerydruginfoResultModel() *AlibabaAlihealthDrugKytQuerydruginfoResultModel {
	return poolAlibabaAlihealthDrugKytQuerydruginfoResultModel.Get().(*AlibabaAlihealthDrugKytQuerydruginfoResultModel)
}

// ReleaseAlibabaAlihealthDrugKytQuerydruginfoResultModel 释放AlibabaAlihealthDrugKytQuerydruginfoResultModel
func ReleaseAlibabaAlihealthDrugKytQuerydruginfoResultModel(v *AlibabaAlihealthDrugKytQuerydruginfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytQuerydruginfoResultModel.Put(v)
}
