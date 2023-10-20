package xhotelitem

import (
	"sync"
)

// TagQueryResult 结构体
type TagQueryResult struct {
	// 列表
	TagEntityDoList []TagEntityDoList `json:"tag_entity_do_list,omitempty" xml:"tag_entity_do_list>tag_entity_do_list,omitempty"`
	// token
	TokenStr string `json:"token_str,omitempty" xml:"token_str,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总数
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 耗时
	SpentTime int64 `json:"spent_time,omitempty" xml:"spent_time,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTagQueryResult = sync.Pool{
	New: func() any {
		return new(TagQueryResult)
	},
}

// GetTagQueryResult() 从对象池中获取TagQueryResult
func GetTagQueryResult() *TagQueryResult {
	return poolTagQueryResult.Get().(*TagQueryResult)
}

// ReleaseTagQueryResult 释放TagQueryResult
func ReleaseTagQueryResult(v *TagQueryResult) {
	v.TagEntityDoList = v.TagEntityDoList[:0]
	v.TokenStr = ""
	v.ErrorMsg = ""
	v.TotalAmount = 0
	v.SpentTime = 0
	v.Success = false
	poolTagQueryResult.Put(v)
}
