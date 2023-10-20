package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel 结构体
type AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel struct {
	// model
	ModelList []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel
func GetAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel() *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel {
	return poolAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel.Get().(*AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel 释放AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel
func ReleaseAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel(v *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel.Put(v)
}
