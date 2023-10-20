package axintrade

import (
	"sync"
)

// TaobaoAlitripAxinTransFundConfirmResult 结构体
type TaobaoAlitripAxinTransFundConfirmResult struct {
	// 描述信息
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口出参
	Data *AxinFundConfirmResDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
}

var poolTaobaoAlitripAxinTransFundConfirmResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransFundConfirmResult)
	},
}

// GetTaobaoAlitripAxinTransFundConfirmResult() 从对象池中获取TaobaoAlitripAxinTransFundConfirmResult
func GetTaobaoAlitripAxinTransFundConfirmResult() *TaobaoAlitripAxinTransFundConfirmResult {
	return poolTaobaoAlitripAxinTransFundConfirmResult.Get().(*TaobaoAlitripAxinTransFundConfirmResult)
}

// ReleaseTaobaoAlitripAxinTransFundConfirmResult 释放TaobaoAlitripAxinTransFundConfirmResult
func ReleaseTaobaoAlitripAxinTransFundConfirmResult(v *TaobaoAlitripAxinTransFundConfirmResult) {
	v.InfoMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	v.NeedRetry = false
	poolTaobaoAlitripAxinTransFundConfirmResult.Put(v)
}
