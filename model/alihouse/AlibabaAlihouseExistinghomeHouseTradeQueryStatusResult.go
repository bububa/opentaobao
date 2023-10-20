package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult 结构体
type AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败成功信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果对象
	Data *SyncHouseStatusDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult
func GetAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult() *AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult {
	return poolAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult.Get().(*AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult 释放AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult
func ReleaseAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult(v *AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeHouseTradeQueryStatusResult.Put(v)
}
