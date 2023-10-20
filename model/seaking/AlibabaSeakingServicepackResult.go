package seaking

import (
	"sync"
)

// AlibabaSeakingServicepackResult 结构体
type AlibabaSeakingServicepackResult struct {
	// 到期时间
	ValidateTo string `json:"validate_to,omitempty" xml:"validate_to,omitempty"`
	// 权限包名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权限包id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolAlibabaSeakingServicepackResult = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingServicepackResult)
	},
}

// GetAlibabaSeakingServicepackResult() 从对象池中获取AlibabaSeakingServicepackResult
func GetAlibabaSeakingServicepackResult() *AlibabaSeakingServicepackResult {
	return poolAlibabaSeakingServicepackResult.Get().(*AlibabaSeakingServicepackResult)
}

// ReleaseAlibabaSeakingServicepackResult 释放AlibabaSeakingServicepackResult
func ReleaseAlibabaSeakingServicepackResult(v *AlibabaSeakingServicepackResult) {
	v.ValidateTo = ""
	v.Name = ""
	v.Id = 0
	poolAlibabaSeakingServicepackResult.Put(v)
}
