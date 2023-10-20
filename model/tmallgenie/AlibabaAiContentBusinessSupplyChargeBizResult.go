package tmallgenie

import (
	"sync"
)

// AlibabaAiContentBusinessSupplyChargeBizResult 结构体
type AlibabaAiContentBusinessSupplyChargeBizResult struct {
	// 0 表示 请求成功，其他表示请求失败
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// true为充值成功
	RetValue bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

var poolAlibabaAiContentBusinessSupplyChargeBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAiContentBusinessSupplyChargeBizResult)
	},
}

// GetAlibabaAiContentBusinessSupplyChargeBizResult() 从对象池中获取AlibabaAiContentBusinessSupplyChargeBizResult
func GetAlibabaAiContentBusinessSupplyChargeBizResult() *AlibabaAiContentBusinessSupplyChargeBizResult {
	return poolAlibabaAiContentBusinessSupplyChargeBizResult.Get().(*AlibabaAiContentBusinessSupplyChargeBizResult)
}

// ReleaseAlibabaAiContentBusinessSupplyChargeBizResult 释放AlibabaAiContentBusinessSupplyChargeBizResult
func ReleaseAlibabaAiContentBusinessSupplyChargeBizResult(v *AlibabaAiContentBusinessSupplyChargeBizResult) {
	v.RetCode = ""
	v.RetMsg = ""
	v.RetValue = false
	poolAlibabaAiContentBusinessSupplyChargeBizResult.Put(v)
}
