package ihome

import (
	"sync"
)

// TaobaoIhomeAdvancepicUploadApiResult 结构体
type TaobaoIhomeAdvancepicUploadApiResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 提交成功的批次id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoIhomeAdvancepicUploadApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoIhomeAdvancepicUploadApiResult)
	},
}

// GetTaobaoIhomeAdvancepicUploadApiResult() 从对象池中获取TaobaoIhomeAdvancepicUploadApiResult
func GetTaobaoIhomeAdvancepicUploadApiResult() *TaobaoIhomeAdvancepicUploadApiResult {
	return poolTaobaoIhomeAdvancepicUploadApiResult.Get().(*TaobaoIhomeAdvancepicUploadApiResult)
}

// ReleaseTaobaoIhomeAdvancepicUploadApiResult 释放TaobaoIhomeAdvancepicUploadApiResult
func ReleaseTaobaoIhomeAdvancepicUploadApiResult(v *TaobaoIhomeAdvancepicUploadApiResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = 0
	v.Success = false
	poolTaobaoIhomeAdvancepicUploadApiResult.Put(v)
}
