package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBrokerPointsSyncResult 结构体
type AlibabaAlihouseExistinghomeBrokerPointsSyncResult struct {
	// 操作结果
	Data []BrokerPointResultDto `json:"data,omitempty" xml:"data>broker_point_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBrokerPointsSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerPointsSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerPointsSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeBrokerPointsSyncResult
func GetAlibabaAlihouseExistinghomeBrokerPointsSyncResult() *AlibabaAlihouseExistinghomeBrokerPointsSyncResult {
	return poolAlibabaAlihouseExistinghomeBrokerPointsSyncResult.Get().(*AlibabaAlihouseExistinghomeBrokerPointsSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerPointsSyncResult 释放AlibabaAlihouseExistinghomeBrokerPointsSyncResult
func ReleaseAlibabaAlihouseExistinghomeBrokerPointsSyncResult(v *AlibabaAlihouseExistinghomeBrokerPointsSyncResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeBrokerPointsSyncResult.Put(v)
}
