package taotv

import (
	"sync"
)

// TaobaoTaotvCarouselCategoryListResult 结构体
type TaobaoTaotvCarouselCategoryListResult struct {
	// 数据列表
	ModelList []TaobaoTaotvCarouselCategoryListModel `json:"model_list,omitempty" xml:"model_list>taobao_taotv_carousel_category_list_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvCarouselCategoryListResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselCategoryListResult)
	},
}

// GetTaobaoTaotvCarouselCategoryListResult() 从对象池中获取TaobaoTaotvCarouselCategoryListResult
func GetTaobaoTaotvCarouselCategoryListResult() *TaobaoTaotvCarouselCategoryListResult {
	return poolTaobaoTaotvCarouselCategoryListResult.Get().(*TaobaoTaotvCarouselCategoryListResult)
}

// ReleaseTaobaoTaotvCarouselCategoryListResult 释放TaobaoTaotvCarouselCategoryListResult
func ReleaseTaobaoTaotvCarouselCategoryListResult(v *TaobaoTaotvCarouselCategoryListResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvCarouselCategoryListResult.Put(v)
}
