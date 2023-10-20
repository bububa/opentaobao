package idleisv

import (
	"sync"
)

// AlibabaIdleIsvPvListResult 结构体
type AlibabaIdleIsvPvListResult struct {
	// 品牌/型号两级属性
	PropertyList []IdleNewPubPropertyValueDo `json:"property_list,omitempty" xml:"property_list>idle_new_pub_property_value_do,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvPvListResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvPvListResult)
	},
}

// GetAlibabaIdleIsvPvListResult() 从对象池中获取AlibabaIdleIsvPvListResult
func GetAlibabaIdleIsvPvListResult() *AlibabaIdleIsvPvListResult {
	return poolAlibabaIdleIsvPvListResult.Get().(*AlibabaIdleIsvPvListResult)
}

// ReleaseAlibabaIdleIsvPvListResult 释放AlibabaIdleIsvPvListResult
func ReleaseAlibabaIdleIsvPvListResult(v *AlibabaIdleIsvPvListResult) {
	v.PropertyList = v.PropertyList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaIdleIsvPvListResult.Put(v)
}
