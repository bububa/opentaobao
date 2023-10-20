package alitrippoi

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 返回的数据实体
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 时间
	ServerTime int64 `json:"server_time,omitempty" xml:"server_time,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.Data = ""
	v.Title = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ServerTime = 0
	v.Success = false
	poolBaseResult.Put(v)
}
