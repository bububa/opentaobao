package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSyncResult 结构体
type AlibabaAlihouseExistinghomeHouseFeaturesSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作返回值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseFeaturesSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseFeaturesSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseFeaturesSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseFeaturesSyncResult
func GetAlibabaAlihouseExistinghomeHouseFeaturesSyncResult() *AlibabaAlihouseExistinghomeHouseFeaturesSyncResult {
	return poolAlibabaAlihouseExistinghomeHouseFeaturesSyncResult.Get().(*AlibabaAlihouseExistinghomeHouseFeaturesSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncResult 释放AlibabaAlihouseExistinghomeHouseFeaturesSyncResult
func ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncResult(v *AlibabaAlihouseExistinghomeHouseFeaturesSyncResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseExistinghomeHouseFeaturesSyncResult.Put(v)
}
