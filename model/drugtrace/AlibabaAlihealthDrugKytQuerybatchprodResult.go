package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytQuerybatchprodResult 结构体
type AlibabaAlihealthDrugKytQuerybatchprodResult struct {
	// 批次产品信息DTO
	Models []BatchProductInfoDto `json:"models,omitempty" xml:"models>batch_product_info_dto,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytQuerybatchprodResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerybatchprodResult)
	},
}

// GetAlibabaAlihealthDrugKytQuerybatchprodResult() 从对象池中获取AlibabaAlihealthDrugKytQuerybatchprodResult
func GetAlibabaAlihealthDrugKytQuerybatchprodResult() *AlibabaAlihealthDrugKytQuerybatchprodResult {
	return poolAlibabaAlihealthDrugKytQuerybatchprodResult.Get().(*AlibabaAlihealthDrugKytQuerybatchprodResult)
}

// ReleaseAlibabaAlihealthDrugKytQuerybatchprodResult 释放AlibabaAlihealthDrugKytQuerybatchprodResult
func ReleaseAlibabaAlihealthDrugKytQuerybatchprodResult(v *AlibabaAlihealthDrugKytQuerybatchprodResult) {
	v.Models = v.Models[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaAlihealthDrugKytQuerybatchprodResult.Put(v)
}
