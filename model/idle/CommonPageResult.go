package idle

import (
	"sync"
)

// CommonPageResult 结构体
type CommonPageResult struct {
	// 商品数据
	Data []TenderItemListVo `json:"data,omitempty" xml:"data>tender_item_list_vo,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 异常信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 数据总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 请求结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否有下页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
}

var poolCommonPageResult = sync.Pool{
	New: func() any {
		return new(CommonPageResult)
	},
}

// GetCommonPageResult() 从对象池中获取CommonPageResult
func GetCommonPageResult() *CommonPageResult {
	return poolCommonPageResult.Get().(*CommonPageResult)
}

// ReleaseCommonPageResult 释放CommonPageResult
func ReleaseCommonPageResult(v *CommonPageResult) {
	v.Data = v.Data[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.CurrentPage = 0
	v.Total = 0
	v.Success = false
	v.HasNextPage = false
	poolCommonPageResult.Put(v)
}
