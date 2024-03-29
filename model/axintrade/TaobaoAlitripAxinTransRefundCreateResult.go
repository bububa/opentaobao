package axintrade

import (
	"sync"
)

// TaobaoAlitripAxinTransRefundCreateResult 结构体
type TaobaoAlitripAxinTransRefundCreateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 提示信息
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 返回素材id
	Data *AxinRefundCreateResDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripAxinTransRefundCreateResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransRefundCreateResult)
	},
}

// GetTaobaoAlitripAxinTransRefundCreateResult() 从对象池中获取TaobaoAlitripAxinTransRefundCreateResult
func GetTaobaoAlitripAxinTransRefundCreateResult() *TaobaoAlitripAxinTransRefundCreateResult {
	return poolTaobaoAlitripAxinTransRefundCreateResult.Get().(*TaobaoAlitripAxinTransRefundCreateResult)
}

// ReleaseTaobaoAlitripAxinTransRefundCreateResult 释放TaobaoAlitripAxinTransRefundCreateResult
func ReleaseTaobaoAlitripAxinTransRefundCreateResult(v *TaobaoAlitripAxinTransRefundCreateResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.InfoMsg = ""
	v.Data = nil
	v.NeedRetry = false
	v.Success = false
	poolTaobaoAlitripAxinTransRefundCreateResult.Put(v)
}
