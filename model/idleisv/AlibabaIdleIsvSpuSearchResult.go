package idleisv

import (
	"sync"
)

// AlibabaIdleIsvSpuSearchResult 结构体
type AlibabaIdleIsvSpuSearchResult struct {
	// 候选的品牌型号列表
	SpuList []SpuPVDo `json:"spu_list,omitempty" xml:"spu_list>spu_pv_do,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvSpuSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvSpuSearchResult)
	},
}

// GetAlibabaIdleIsvSpuSearchResult() 从对象池中获取AlibabaIdleIsvSpuSearchResult
func GetAlibabaIdleIsvSpuSearchResult() *AlibabaIdleIsvSpuSearchResult {
	return poolAlibabaIdleIsvSpuSearchResult.Get().(*AlibabaIdleIsvSpuSearchResult)
}

// ReleaseAlibabaIdleIsvSpuSearchResult 释放AlibabaIdleIsvSpuSearchResult
func ReleaseAlibabaIdleIsvSpuSearchResult(v *AlibabaIdleIsvSpuSearchResult) {
	v.SpuList = v.SpuList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaIdleIsvSpuSearchResult.Put(v)
}
