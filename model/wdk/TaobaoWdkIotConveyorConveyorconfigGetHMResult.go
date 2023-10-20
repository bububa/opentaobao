package wdk

import (
	"sync"
)

// TaobaoWdkIotConveyorConveyorconfigGetHMResult 结构体
type TaobaoWdkIotConveyorConveyorconfigGetHMResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 配置信息列表
	Model *ConveyorBasicConfigDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWdkIotConveyorConveyorconfigGetHMResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkIotConveyorConveyorconfigGetHMResult)
	},
}

// GetTaobaoWdkIotConveyorConveyorconfigGetHMResult() 从对象池中获取TaobaoWdkIotConveyorConveyorconfigGetHMResult
func GetTaobaoWdkIotConveyorConveyorconfigGetHMResult() *TaobaoWdkIotConveyorConveyorconfigGetHMResult {
	return poolTaobaoWdkIotConveyorConveyorconfigGetHMResult.Get().(*TaobaoWdkIotConveyorConveyorconfigGetHMResult)
}

// ReleaseTaobaoWdkIotConveyorConveyorconfigGetHMResult 释放TaobaoWdkIotConveyorConveyorconfigGetHMResult
func ReleaseTaobaoWdkIotConveyorConveyorconfigGetHMResult(v *TaobaoWdkIotConveyorConveyorconfigGetHMResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Model = nil
	v.Success = false
	poolTaobaoWdkIotConveyorConveyorconfigGetHMResult.Put(v)
}
