package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBrandcitySyncResult 结构体
type AlibabaAlihouseExistinghomeBrandcitySyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBrandcitySyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrandcitySyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeBrandcitySyncResult() 从对象池中获取AlibabaAlihouseExistinghomeBrandcitySyncResult
func GetAlibabaAlihouseExistinghomeBrandcitySyncResult() *AlibabaAlihouseExistinghomeBrandcitySyncResult {
	return poolAlibabaAlihouseExistinghomeBrandcitySyncResult.Get().(*AlibabaAlihouseExistinghomeBrandcitySyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeBrandcitySyncResult 释放AlibabaAlihouseExistinghomeBrandcitySyncResult
func ReleaseAlibabaAlihouseExistinghomeBrandcitySyncResult(v *AlibabaAlihouseExistinghomeBrandcitySyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeBrandcitySyncResult.Put(v)
}
