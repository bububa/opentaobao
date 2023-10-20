package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQueryUpbillcodeResultModel 结构体
type AlibabaAlihealthDrugKytQueryUpbillcodeResultModel struct {
	// model
	ModelList []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQueryUpbillcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryUpbillcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytQueryUpbillcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytQueryUpbillcodeResultModel
func GetAlibabaAlihealthDrugKytQueryUpbillcodeResultModel() *AlibabaAlihealthDrugKytQueryUpbillcodeResultModel {
	return poolAlibabaAlihealthDrugKytQueryUpbillcodeResultModel.Get().(*AlibabaAlihealthDrugKytQueryUpbillcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeResultModel 释放AlibabaAlihealthDrugKytQueryUpbillcodeResultModel
func ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeResultModel(v *AlibabaAlihealthDrugKytQueryUpbillcodeResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugKytQueryUpbillcodeResultModel.Put(v)
}
