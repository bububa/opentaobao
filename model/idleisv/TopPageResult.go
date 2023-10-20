package idleisv

import (
	"sync"
)

// TopPageResult 结构体
type TopPageResult struct {
	// 商品列表
	ItemList []IdleItemApiDo `json:"item_list,omitempty" xml:"item_list>idle_item_api_do,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 总结果数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否有下一页
	NextPage bool `json:"next_page,omitempty" xml:"next_page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopPageResult = sync.Pool{
	New: func() any {
		return new(TopPageResult)
	},
}

// GetTopPageResult() 从对象池中获取TopPageResult
func GetTopPageResult() *TopPageResult {
	return poolTopPageResult.Get().(*TopPageResult)
}

// ReleaseTopPageResult 释放TopPageResult
func ReleaseTopPageResult(v *TopPageResult) {
	v.ItemList = v.ItemList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Total = 0
	v.NextPage = false
	v.Success = false
	poolTopPageResult.Put(v)
}
