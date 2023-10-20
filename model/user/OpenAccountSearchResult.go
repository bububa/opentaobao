package user

import (
	"sync"
)

// OpenAccountSearchResult 结构体
type OpenAccountSearchResult struct {
	// OpenAccount的列表
	Datas []OpenAccount `json:"datas,omitempty" xml:"datas>open_account,omitempty"`
	// 状态信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 查询是否成功，成功返回时有可能数据为空
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

var poolOpenAccountSearchResult = sync.Pool{
	New: func() any {
		return new(OpenAccountSearchResult)
	},
}

// GetOpenAccountSearchResult() 从对象池中获取OpenAccountSearchResult
func GetOpenAccountSearchResult() *OpenAccountSearchResult {
	return poolOpenAccountSearchResult.Get().(*OpenAccountSearchResult)
}

// ReleaseOpenAccountSearchResult 释放OpenAccountSearchResult
func ReleaseOpenAccountSearchResult(v *OpenAccountSearchResult) {
	v.Datas = v.Datas[:0]
	v.Message = ""
	v.Code = 0
	v.Successful = false
	poolOpenAccountSearchResult.Put(v)
}
