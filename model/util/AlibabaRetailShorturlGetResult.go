package util

import (
	"sync"
)

// AlibabaRetailShorturlGetResult 结构体
type AlibabaRetailShorturlGetResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// module
	Module *ShortUrlDto `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailShorturlGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailShorturlGetResult)
	},
}

// GetAlibabaRetailShorturlGetResult() 从对象池中获取AlibabaRetailShorturlGetResult
func GetAlibabaRetailShorturlGetResult() *AlibabaRetailShorturlGetResult {
	return poolAlibabaRetailShorturlGetResult.Get().(*AlibabaRetailShorturlGetResult)
}

// ReleaseAlibabaRetailShorturlGetResult 释放AlibabaRetailShorturlGetResult
func ReleaseAlibabaRetailShorturlGetResult(v *AlibabaRetailShorturlGetResult) {
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Module = nil
	v.Success = false
	poolAlibabaRetailShorturlGetResult.Put(v)
}
