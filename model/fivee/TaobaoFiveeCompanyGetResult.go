package fivee

import (
	"sync"
)

// TaobaoFiveeCompanyGetResult 结构体
type TaobaoFiveeCompanyGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *Company `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFiveeCompanyGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeCompanyGetResult)
	},
}

// GetTaobaoFiveeCompanyGetResult() 从对象池中获取TaobaoFiveeCompanyGetResult
func GetTaobaoFiveeCompanyGetResult() *TaobaoFiveeCompanyGetResult {
	return poolTaobaoFiveeCompanyGetResult.Get().(*TaobaoFiveeCompanyGetResult)
}

// ReleaseTaobaoFiveeCompanyGetResult 释放TaobaoFiveeCompanyGetResult
func ReleaseTaobaoFiveeCompanyGetResult(v *TaobaoFiveeCompanyGetResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	v.Success = false
	poolTaobaoFiveeCompanyGetResult.Put(v)
}
