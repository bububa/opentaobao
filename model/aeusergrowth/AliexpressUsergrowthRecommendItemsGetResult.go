package aeusergrowth

import (
	"sync"
)

// AliexpressUsergrowthRecommendItemsGetResult 结构体
type AliexpressUsergrowthRecommendItemsGetResult struct {
	// Result itemList,The product are located at the top,maybe null when success = false
	DataList []AliexpressUsergrowthRecommendItemsGetData `json:"data_list,omitempty" xml:"data_list>aliexpress_usergrowth_recommend_items_get_data,omitempty"`
	// other extend message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// Code is used to determine whether the result is correct
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// success is used to determine whether invoke service success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressUsergrowthRecommendItemsGetResult = sync.Pool{
	New: func() any {
		return new(AliexpressUsergrowthRecommendItemsGetResult)
	},
}

// GetAliexpressUsergrowthRecommendItemsGetResult() 从对象池中获取AliexpressUsergrowthRecommendItemsGetResult
func GetAliexpressUsergrowthRecommendItemsGetResult() *AliexpressUsergrowthRecommendItemsGetResult {
	return poolAliexpressUsergrowthRecommendItemsGetResult.Get().(*AliexpressUsergrowthRecommendItemsGetResult)
}

// ReleaseAliexpressUsergrowthRecommendItemsGetResult 释放AliexpressUsergrowthRecommendItemsGetResult
func ReleaseAliexpressUsergrowthRecommendItemsGetResult(v *AliexpressUsergrowthRecommendItemsGetResult) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAliexpressUsergrowthRecommendItemsGetResult.Put(v)
}
