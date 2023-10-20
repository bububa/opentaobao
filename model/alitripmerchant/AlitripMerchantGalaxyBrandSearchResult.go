package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyBrandSearchResult 结构体
type AlitripMerchantGalaxyBrandSearchResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 品牌信息
	Brands *Content `json:"brands,omitempty" xml:"brands,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyBrandSearchResult = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyBrandSearchResult)
	},
}

// GetAlitripMerchantGalaxyBrandSearchResult() 从对象池中获取AlitripMerchantGalaxyBrandSearchResult
func GetAlitripMerchantGalaxyBrandSearchResult() *AlitripMerchantGalaxyBrandSearchResult {
	return poolAlitripMerchantGalaxyBrandSearchResult.Get().(*AlitripMerchantGalaxyBrandSearchResult)
}

// ReleaseAlitripMerchantGalaxyBrandSearchResult 释放AlitripMerchantGalaxyBrandSearchResult
func ReleaseAlitripMerchantGalaxyBrandSearchResult(v *AlitripMerchantGalaxyBrandSearchResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Brands = nil
	v.Success = false
	poolAlitripMerchantGalaxyBrandSearchResult.Put(v)
}
