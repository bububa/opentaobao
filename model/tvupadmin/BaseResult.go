package tvupadmin

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// retValue
	RetValues []SimpleBotInfo `json:"ret_values,omitempty" xml:"ret_values>simple_bot_info,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 返回数据
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 返回码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.RetValues = v.RetValues[:0]
	v.RetMsg = ""
	v.RetValue = ""
	v.RetCode = 0
	poolBaseResult.Put(v)
}
