package omniorder

import (
	"sync"
)

// TaobaoOmniitemCategoryGetResult 结构体
type TaobaoOmniitemCategoryGetResult struct {
	// data
	Datas []OmniItemCategoryDto `json:"datas,omitempty" xml:"datas>omni_item_category_dto,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOmniitemCategoryGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemCategoryGetResult)
	},
}

// GetTaobaoOmniitemCategoryGetResult() 从对象池中获取TaobaoOmniitemCategoryGetResult
func GetTaobaoOmniitemCategoryGetResult() *TaobaoOmniitemCategoryGetResult {
	return poolTaobaoOmniitemCategoryGetResult.Get().(*TaobaoOmniitemCategoryGetResult)
}

// ReleaseTaobaoOmniitemCategoryGetResult 释放TaobaoOmniitemCategoryGetResult
func ReleaseTaobaoOmniitemCategoryGetResult(v *TaobaoOmniitemCategoryGetResult) {
	v.Datas = v.Datas[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolTaobaoOmniitemCategoryGetResult.Put(v)
}
