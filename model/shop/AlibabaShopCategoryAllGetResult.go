package shop

import (
	"sync"
)

// AlibabaShopCategoryAllGetResult 结构体
type AlibabaShopCategoryAllGetResult struct {
	// 分类对象
	ModuleList []OpenCategoryDto `json:"module_list,omitempty" xml:"module_list>open_category_dto,omitempty"`
	// 返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分类总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaShopCategoryAllGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaShopCategoryAllGetResult)
	},
}

// GetAlibabaShopCategoryAllGetResult() 从对象池中获取AlibabaShopCategoryAllGetResult
func GetAlibabaShopCategoryAllGetResult() *AlibabaShopCategoryAllGetResult {
	return poolAlibabaShopCategoryAllGetResult.Get().(*AlibabaShopCategoryAllGetResult)
}

// ReleaseAlibabaShopCategoryAllGetResult 释放AlibabaShopCategoryAllGetResult
func ReleaseAlibabaShopCategoryAllGetResult(v *AlibabaShopCategoryAllGetResult) {
	v.ModuleList = v.ModuleList[:0]
	v.Message = ""
	v.Total = 0
	v.Success = false
	poolAlibabaShopCategoryAllGetResult.Put(v)
}
