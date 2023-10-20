package wdk

import (
	"sync"
)

// AlibabaWdkShopQueryApiResults 结构体
type AlibabaWdkShopQueryApiResults struct {
	// 返回门店信息列表
	Models []ShopDo `json:"models,omitempty" xml:"models>shop_do,omitempty"`
	// 异常编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// true
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkShopQueryApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkShopQueryApiResults)
	},
}

// GetAlibabaWdkShopQueryApiResults() 从对象池中获取AlibabaWdkShopQueryApiResults
func GetAlibabaWdkShopQueryApiResults() *AlibabaWdkShopQueryApiResults {
	return poolAlibabaWdkShopQueryApiResults.Get().(*AlibabaWdkShopQueryApiResults)
}

// ReleaseAlibabaWdkShopQueryApiResults 释放AlibabaWdkShopQueryApiResults
func ReleaseAlibabaWdkShopQueryApiResults(v *AlibabaWdkShopQueryApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = ""
	poolAlibabaWdkShopQueryApiResults.Put(v)
}
