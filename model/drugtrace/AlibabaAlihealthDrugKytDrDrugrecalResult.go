package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrDrugrecalResult 结构体
type AlibabaAlihealthDrugKytDrDrugrecalResult struct {
	// 召回信息列表
	ModelList []AlibabaAlihealthDrugKytDrDrugrecalModel `json:"model_list,omitempty" xml:"model_list>alibaba_alihealth_drug_kyt_dr_drugrecal_model,omitempty"`
	// 服务返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 服务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrDrugrecalResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrDrugrecalResult)
	},
}

// GetAlibabaAlihealthDrugKytDrDrugrecalResult() 从对象池中获取AlibabaAlihealthDrugKytDrDrugrecalResult
func GetAlibabaAlihealthDrugKytDrDrugrecalResult() *AlibabaAlihealthDrugKytDrDrugrecalResult {
	return poolAlibabaAlihealthDrugKytDrDrugrecalResult.Get().(*AlibabaAlihealthDrugKytDrDrugrecalResult)
}

// ReleaseAlibabaAlihealthDrugKytDrDrugrecalResult 释放AlibabaAlihealthDrugKytDrDrugrecalResult
func ReleaseAlibabaAlihealthDrugKytDrDrugrecalResult(v *AlibabaAlihealthDrugKytDrDrugrecalResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.IsSuccess = false
	poolAlibabaAlihealthDrugKytDrDrugrecalResult.Put(v)
}
