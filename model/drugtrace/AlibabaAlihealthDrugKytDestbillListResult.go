package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDestbillListResult 结构体
type AlibabaAlihealthDrugKytDestbillListResult struct {
	// 接口返回model
	ModelList []AlibabaAlihealthDrugKytDestbillListModel `json:"model_list,omitempty" xml:"model_list>alibaba_alihealth_drug_kyt_destbill_list_model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDestbillListResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDestbillListResult)
	},
}

// GetAlibabaAlihealthDrugKytDestbillListResult() 从对象池中获取AlibabaAlihealthDrugKytDestbillListResult
func GetAlibabaAlihealthDrugKytDestbillListResult() *AlibabaAlihealthDrugKytDestbillListResult {
	return poolAlibabaAlihealthDrugKytDestbillListResult.Get().(*AlibabaAlihealthDrugKytDestbillListResult)
}

// ReleaseAlibabaAlihealthDrugKytDestbillListResult 释放AlibabaAlihealthDrugKytDestbillListResult
func ReleaseAlibabaAlihealthDrugKytDestbillListResult(v *AlibabaAlihealthDrugKytDestbillListResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ResponseSuccess = false
	poolAlibabaAlihealthDrugKytDestbillListResult.Put(v)
}
