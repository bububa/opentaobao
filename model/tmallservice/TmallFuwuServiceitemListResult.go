package tmallservice

import (
	"sync"
)

// TmallFuwuServiceitemListResult 结构体
type TmallFuwuServiceitemListResult struct {
	// 服务商品信息列表的json对象
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorType
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTmallFuwuServiceitemListResult = sync.Pool{
	New: func() any {
		return new(TmallFuwuServiceitemListResult)
	},
}

// GetTmallFuwuServiceitemListResult() 从对象池中获取TmallFuwuServiceitemListResult
func GetTmallFuwuServiceitemListResult() *TmallFuwuServiceitemListResult {
	return poolTmallFuwuServiceitemListResult.Get().(*TmallFuwuServiceitemListResult)
}

// ReleaseTmallFuwuServiceitemListResult 释放TmallFuwuServiceitemListResult
func ReleaseTmallFuwuServiceitemListResult(v *TmallFuwuServiceitemListResult) {
	v.ResultData = ""
	v.Message = ""
	v.ErrorCode = ""
	v.ErrorType = 0
	v.IsSuccess = false
	poolTmallFuwuServiceitemListResult.Put(v)
}
